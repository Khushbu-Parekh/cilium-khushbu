// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package observer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"net"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"

	flowpb "github.com/cilium/cilium/api/v1/flow"
	observerpb "github.com/cilium/cilium/api/v1/observer"
	relaypb "github.com/cilium/cilium/api/v1/relay"
	"github.com/cilium/cilium/pkg/hubble/defaults"
	peerTypes "github.com/cilium/cilium/pkg/hubble/peer/types"
	poolTypes "github.com/cilium/cilium/pkg/hubble/relay/pool/types"
	"github.com/cilium/cilium/pkg/hubble/testutils"
	"github.com/cilium/cilium/pkg/logging"
)

func TestGetFlows(t *testing.T) {
	type results struct {
		numFlows     int
		flows        map[string][]*flowpb.Flow
		statusEvents []*relaypb.NodeStatusEvent
	}
	var got *results
	type want struct {
		flows        map[string][]*flowpb.Flow
		statusEvents []*relaypb.NodeStatusEvent
		err          error
		log          []string
	}
	fss := &testutils.FakeGRPCServerStream{
		OnContext: context.TODO,
	}
	done := make(chan struct{})
	tests := []struct {
		name   string
		plr    PeerLister
		ocb    observerClientBuilder
		req    *observerpb.GetFlowsRequest
		stream observerpb.Observer_GetFlowsServer
		want   want
	}{
		{
			name: "Observe 0 flows from 1 peer without address",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name:    "noip",
								Address: nil,
							},
							Conn: nil,
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{},
			req: &observerpb.GetFlowsRequest{Number: 0},
			stream: &testutils.FakeGetFlowsServer{
				FakeGRPCServerStream: fss,
				OnSend: func(resp *observerpb.GetFlowsResponse) error {
					if resp == nil {
						return nil
					}
					switch resp.GetResponseTypes().(type) {
					case *observerpb.GetFlowsResponse_Flow:
						got.numFlows++
						got.flows[resp.GetNodeName()] = append(got.flows[resp.GetNodeName()], resp.GetFlow())
					case *observerpb.GetFlowsResponse_NodeStatus:
						got.statusEvents = append(got.statusEvents, resp.GetNodeStatus())
					}
					if got.numFlows == 0 && len(got.statusEvents) == 1 {
						close(done)
						return io.EOF
					}
					return nil
				},
			},
			want: want{
				flows: map[string][]*flowpb.Flow{},
				statusEvents: []*relaypb.NodeStatusEvent{
					{
						StateChange: relaypb.NodeState_NODE_UNAVAILABLE,
						NodeNames:   []string{"noip"},
					},
				},
				err: io.EOF,
			},
		}, {
			name: "Observe 4 flows from 2 online peers",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					var numRecv uint64
					return &testutils.FakeObserverClient{
						OnGetFlows: func(_ context.Context, in *observerpb.GetFlowsRequest, _ ...grpc.CallOption) (observerpb.Observer_GetFlowsClient, error) {
							return &testutils.FakeGetFlowsClient{
								OnRecv: func() (*observerpb.GetFlowsResponse, error) {
									if numRecv == in.Number {
										return nil, io.EOF
									}
									numRecv++
									return &observerpb.GetFlowsResponse{
										NodeName: p.Name,
										ResponseTypes: &observerpb.GetFlowsResponse_Flow{
											Flow: &flowpb.Flow{
												NodeName: p.Name,
											},
										},
									}, nil
								},
							}, nil
						},
					}
				},
			},
			req: &observerpb.GetFlowsRequest{Number: 2},
			stream: &testutils.FakeGetFlowsServer{
				FakeGRPCServerStream: fss,
				OnSend: func(resp *observerpb.GetFlowsResponse) error {
					if resp == nil {
						return nil
					}
					switch resp.GetResponseTypes().(type) {
					case *observerpb.GetFlowsResponse_Flow:
						got.numFlows++
						got.flows[resp.GetNodeName()] = append(got.flows[resp.GetNodeName()], resp.GetFlow())
					case *observerpb.GetFlowsResponse_NodeStatus:
						got.statusEvents = append(got.statusEvents, resp.GetNodeStatus())
					}
					if got.numFlows == 4 && len(got.statusEvents) == 1 {
						close(done)
						return io.EOF
					}
					return nil
				},
			},
			want: want{
				flows: map[string][]*flowpb.Flow{
					"one": {&flowpb.Flow{NodeName: "one"}, &flowpb.Flow{NodeName: "one"}},
					"two": {&flowpb.Flow{NodeName: "two"}, &flowpb.Flow{NodeName: "two"}},
				},
				statusEvents: []*relaypb.NodeStatusEvent{
					{
						StateChange: relaypb.NodeState_NODE_CONNECTED,
						NodeNames:   []string{"one", "two"},
					},
				},
				err: io.EOF,
			},
		}, {
			name: "Observe 2 flows from 1 online peer and none from 1 unavailable peer",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					var numRecv uint64
					return &testutils.FakeObserverClient{
						OnGetFlows: func(_ context.Context, in *observerpb.GetFlowsRequest, _ ...grpc.CallOption) (observerpb.Observer_GetFlowsClient, error) {
							if p.Name != "one" {
								return nil, fmt.Errorf("GetFlows() called for peer '%s'; this is unexpected", p.Name)
							}
							return &testutils.FakeGetFlowsClient{
								OnRecv: func() (*observerpb.GetFlowsResponse, error) {
									if numRecv == in.Number {
										return nil, io.EOF
									}
									numRecv++
									return &observerpb.GetFlowsResponse{
										NodeName: p.Name,
										ResponseTypes: &observerpb.GetFlowsResponse_Flow{
											Flow: &flowpb.Flow{
												NodeName: p.Name,
											},
										},
									}, nil
								},
							}, nil
						},
					}
				},
			},
			req: &observerpb.GetFlowsRequest{Number: 2},
			stream: &testutils.FakeGetFlowsServer{
				FakeGRPCServerStream: fss,
				OnSend: func(resp *observerpb.GetFlowsResponse) error {
					if resp == nil {
						return nil
					}
					switch resp.GetResponseTypes().(type) {
					case *observerpb.GetFlowsResponse_Flow:
						got.numFlows++
						got.flows[resp.GetNodeName()] = append(got.flows[resp.GetNodeName()], resp.GetFlow())
					case *observerpb.GetFlowsResponse_NodeStatus:
						got.statusEvents = append(got.statusEvents, resp.GetNodeStatus())
					}
					if got.numFlows == 2 && len(got.statusEvents) == 2 {
						close(done)
						return io.EOF
					}
					return nil
				},
			},
			want: want{
				flows: map[string][]*flowpb.Flow{
					"one": {&flowpb.Flow{NodeName: "one"}, &flowpb.Flow{NodeName: "one"}},
				},
				statusEvents: []*relaypb.NodeStatusEvent{
					{
						StateChange: relaypb.NodeState_NODE_CONNECTED,
						NodeNames:   []string{"one"},
					}, {
						StateChange: relaypb.NodeState_NODE_UNAVAILABLE,
						NodeNames:   []string{"two"},
					},
				},
				err: io.EOF,
				log: []string{
					`level=info msg="No connection to peer, skipping" address=192.0.2.2:4244 peer=two`,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got = &results{
				flows: make(map[string][]*flowpb.Flow),
			}
			done = make(chan struct{})
			var buf bytes.Buffer
			logger := slog.New(
				slog.NewTextHandler(&buf,
					&slog.HandlerOptions{
						ReplaceAttr: logging.ReplaceAttrFnWithoutTimestamp,
					},
				),
			)

			srv, err := NewServer(
				tt.plr,
				WithLogger(logger),
				withObserverClientBuilder(tt.ocb),
			)
			assert.NoError(t, err)
			err = srv.GetFlows(tt.req, tt.stream)
			<-done
			assert.Equal(t, tt.want.err, err)
			if diff := cmp.Diff(tt.want.flows, got.flows, cmpopts.IgnoreUnexported(flowpb.Flow{})); diff != "" {
				t.Errorf("Flows mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want.statusEvents, got.statusEvents, cmpopts.IgnoreUnexported(relaypb.NodeStatusEvent{})); diff != "" {
				t.Errorf("StatusEvents mismatch (-want +got):\n%s", diff)
			}
			out := buf.String()
			for _, msg := range tt.want.log {
				assert.Contains(t, out, msg)
			}
		})
	}
}

// test that Relay pick up a joining Hubble peer.
func TestGetFlows_follow(t *testing.T) {
	plChan := make(chan []poolTypes.Peer, 1)
	pl := &testutils.FakePeerLister{
		OnList: func() []poolTypes.Peer {
			return <-plChan
		},
	}
	type resp struct {
		resp *observerpb.GetFlowsResponse
		err  error
	}
	oneChan := make(chan resp, 1)
	one := poolTypes.Peer{
		Peer: peerTypes.Peer{
			Name: "one",
			Address: &net.TCPAddr{
				IP:   net.ParseIP("192.0.2.1"),
				Port: defaults.ServerPort,
			},
		},
		Conn: &testutils.FakeClientConn{
			OnGetState: func() connectivity.State {
				return connectivity.Ready
			},
		},
	}
	twoChan := make(chan resp, 1)
	two := poolTypes.Peer{
		Peer: peerTypes.Peer{
			Name: "two",
			Address: &net.TCPAddr{
				IP:   net.ParseIP("192.0.2.2"),
				Port: defaults.ServerPort,
			},
		},
		Conn: &testutils.FakeClientConn{
			OnGetState: func() connectivity.State {
				return connectivity.Ready
			},
		},
	}

	ocb := fakeObserverClientBuilder{
		onObserverClient: func(peer *poolTypes.Peer) observerpb.ObserverClient {
			return &testutils.FakeObserverClient{
				OnGetFlows: func(_ context.Context, in *observerpb.GetFlowsRequest, _ ...grpc.CallOption) (observerpb.Observer_GetFlowsClient, error) {
					return &testutils.FakeGetFlowsClient{
						OnRecv: func() (*observerpb.GetFlowsResponse, error) {
							switch peer.Name {
							case "one":
								r := <-oneChan
								return r.resp, r.err
							case "two":
								r := <-twoChan
								return r.resp, r.err
							}
							return nil, fmt.Errorf("unexpected peer %q", peer.Name)
						},
					}, nil
				},
			}
		},
	}
	fss := &testutils.FakeGRPCServerStream{
		OnContext: context.TODO,
	}
	seenOneFlows := atomic.Int64{}
	seenTwoFlows := atomic.Int64{}
	stream := &testutils.FakeGetFlowsServer{
		FakeGRPCServerStream: fss,
		OnSend: func(resp *observerpb.GetFlowsResponse) error {
			if resp == nil {
				return nil
			}
			switch resp.GetResponseTypes().(type) {
			case *observerpb.GetFlowsResponse_Flow:
				switch resp.NodeName {
				case "one":
					seenOneFlows.Add(1)
				case "two":
					seenTwoFlows.Add(1)
				}
			case *observerpb.GetFlowsResponse_NodeStatus:
			}
			return nil
		},
	}
	srv, err := NewServer(
		pl,
		withObserverClientBuilder(ocb),
	)
	srv.opts.peerUpdateInterval = 10 * time.Millisecond
	require.NoError(t, err)

	plChan <- []poolTypes.Peer{one}
	oneChan <- resp{
		resp: &observerpb.GetFlowsResponse{
			NodeName: "one",
			ResponseTypes: &observerpb.GetFlowsResponse_Flow{
				Flow: &flowpb.Flow{
					NodeName: "one",
				},
			},
		},
	}
	go func() {
		err = srv.GetFlows(&observerpb.GetFlowsRequest{Follow: true}, stream)
		assert.NoError(t, err)
	}()
	assert.Eventually(t, func() bool {
		return seenOneFlows.Load() == 1
	}, 10*time.Second, 10*time.Millisecond)

	plChan <- []poolTypes.Peer{one, two}
	oneChan <- resp{
		resp: &observerpb.GetFlowsResponse{
			NodeName: "one",
			ResponseTypes: &observerpb.GetFlowsResponse_Flow{
				Flow: &flowpb.Flow{
					NodeName: "one",
				},
			},
		},
	}
	twoChan <- resp{
		resp: &observerpb.GetFlowsResponse{
			NodeName: "two",
			ResponseTypes: &observerpb.GetFlowsResponse_Flow{
				Flow: &flowpb.Flow{
					NodeName: "two",
				},
			},
		},
	}
	assert.Eventually(t, func() bool {
		return seenOneFlows.Load() == 2 && seenTwoFlows.Load() == 1
	}, 10*time.Second, 10*time.Millisecond)

}

func TestGetNodes(t *testing.T) {
	type want struct {
		resp *observerpb.GetNodesResponse
		err  error
		log  []string
	}
	tests := []struct {
		name string
		plr  PeerLister
		ocb  observerClientBuilder
		req  *observerpb.GetNodesRequest
		want want
	}{
		{
			name: "1 peer without address",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name:    "noip",
								Address: nil,
							},
							Conn: nil,
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{},
			want: want{
				resp: &observerpb.GetNodesResponse{
					Nodes: []*observerpb.Node{
						{
							Name:    "noip",
							Version: "",
							Address: "",
							State:   relaypb.NodeState_NODE_UNAVAILABLE,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						},
					},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=<nil> peer=noip`,
				},
			},
		}, {
			name: "2 connected peers",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.ServerStatusResponse{
									UptimeNs:  123456,
									Version:   "cilium v1.9.0",
									MaxFlows:  4095,
									NumFlows:  4095,
									SeenFlows: 11000,
								}, nil
							case "two":
								return &observerpb.ServerStatusResponse{
									UptimeNs:  555555,
									Version:   "cilium v1.9.0",
									MaxFlows:  2047,
									NumFlows:  2020,
									SeenFlows: 12000,
								}, nil
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.GetNodesResponse{
					Nodes: []*observerpb.Node{
						{
							Name:      "one",
							Version:   "cilium v1.9.0",
							Address:   "192.0.2.1:4244",
							State:     relaypb.NodeState_NODE_CONNECTED,
							UptimeNs:  123456,
							MaxFlows:  4095,
							NumFlows:  4095,
							SeenFlows: 11000,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						}, {
							Name:      "two",
							Version:   "cilium v1.9.0",
							Address:   "192.0.2.2:4244",
							State:     relaypb.NodeState_NODE_CONNECTED,
							UptimeNs:  555555,
							MaxFlows:  2047,
							NumFlows:  2020,
							SeenFlows: 12000,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						},
					},
				},
			},
		}, {
			name: "2 connected peers with TLS",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
								TLSEnabled:    true,
								TLSServerName: "one.default.hubble-grpc.cilium.io",
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
								TLSEnabled:    true,
								TLSServerName: "two.default.hubble-grpc.cilium.io",
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.ServerStatusResponse{
									UptimeNs:  123456,
									Version:   "cilium v1.9.0",
									MaxFlows:  4095,
									NumFlows:  4095,
									SeenFlows: 11000,
								}, nil
							case "two":
								return &observerpb.ServerStatusResponse{
									UptimeNs:  555555,
									Version:   "cilium v1.9.0",
									MaxFlows:  2047,
									NumFlows:  2020,
									SeenFlows: 12000,
								}, nil
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.GetNodesResponse{
					Nodes: []*observerpb.Node{
						{
							Name:      "one",
							Version:   "cilium v1.9.0",
							Address:   "192.0.2.1:4244",
							State:     relaypb.NodeState_NODE_CONNECTED,
							UptimeNs:  123456,
							MaxFlows:  4095,
							NumFlows:  4095,
							SeenFlows: 11000,
							Tls: &observerpb.TLS{
								Enabled:    true,
								ServerName: "one.default.hubble-grpc.cilium.io",
							},
						}, {
							Name:      "two",
							Version:   "cilium v1.9.0",
							Address:   "192.0.2.2:4244",
							State:     relaypb.NodeState_NODE_CONNECTED,
							UptimeNs:  555555,
							MaxFlows:  2047,
							NumFlows:  2020,
							SeenFlows: 12000,
							Tls: &observerpb.TLS{
								Enabled:    true,
								ServerName: "two.default.hubble-grpc.cilium.io",
							},
						},
					},
				},
			},
		}, {
			name: "1 connected peer, 1 unreachable peer",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.ServerStatusResponse{
									UptimeNs:  123456,
									Version:   "cilium v1.9.0",
									MaxFlows:  4095,
									NumFlows:  4095,
									SeenFlows: 11000,
								}, nil
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.GetNodesResponse{
					Nodes: []*observerpb.Node{
						{
							Name:      "one",
							Version:   "cilium v1.9.0",
							Address:   "192.0.2.1:4244",
							State:     relaypb.NodeState_NODE_CONNECTED,
							UptimeNs:  123456,
							MaxFlows:  4095,
							NumFlows:  4095,
							SeenFlows: 11000,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						}, {
							Name:     "two",
							Version:  "",
							Address:  "192.0.2.2:4244",
							State:    relaypb.NodeState_NODE_UNAVAILABLE,
							UptimeNs: 0,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						},
					},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=192.0.2.2:4244 peer=two`,
				},
			},
		}, {
			name: "1 connected peer, 1 unreachable peer, 1 peer with error",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "three",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.3"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.ServerStatusResponse{
									UptimeNs:  123456,
									Version:   "cilium v1.9.0",
									MaxFlows:  4095,
									NumFlows:  4095,
									SeenFlows: 11000,
								}, nil
							case "three":
								return nil, status.Errorf(codes.Unimplemented, "ServerStatus not implemented")
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.GetNodesResponse{
					Nodes: []*observerpb.Node{
						{
							Name:      "one",
							Version:   "cilium v1.9.0",
							Address:   "192.0.2.1:4244",
							UptimeNs:  123456,
							MaxFlows:  4095,
							NumFlows:  4095,
							SeenFlows: 11000,
							State:     relaypb.NodeState_NODE_CONNECTED,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						}, {
							Name:    "two",
							Version: "",
							Address: "192.0.2.2:4244",
							State:   relaypb.NodeState_NODE_UNAVAILABLE,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						}, {
							Name:    "three",
							Version: "",
							Address: "192.0.2.3:4244",
							State:   relaypb.NodeState_NODE_ERROR,
							Tls: &observerpb.TLS{
								Enabled:    false,
								ServerName: "",
							},
						},
					},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=192.0.2.2:4244 peer=two`,
					`level=warn msg="Failed to retrieve server status" error="rpc error: code = Unimplemented desc = ServerStatus not implemented" peer=three`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			logger := slog.New(
				slog.NewTextHandler(&buf,
					&slog.HandlerOptions{
						ReplaceAttr: logging.ReplaceAttrFnWithoutTimestamp,
					},
				),
			)

			srv, err := NewServer(
				tt.plr,
				WithLogger(logger),
				withObserverClientBuilder(tt.ocb),
			)
			assert.NoError(t, err)
			got, err := srv.GetNodes(t.Context(), tt.req)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.resp, got)
			out := buf.String()
			for _, msg := range tt.want.log {
				assert.Contains(t, out, msg)
			}
		})
	}
}

func TestGetNamespaces(t *testing.T) {
	type want struct {
		resp *observerpb.GetNamespacesResponse
		err  error
		log  []string
	}
	tests := []struct {
		name string
		plr  PeerLister
		ocb  observerClientBuilder
		req  *observerpb.GetNamespacesRequest
		want want
	}{
		{
			name: "get no namespaces from 1 peer without address",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name:    "noip",
								Address: nil,
							},
							Conn: nil,
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnGetNamespaces: func(_ context.Context, in *observerpb.GetNamespacesRequest, _ ...grpc.CallOption) (*observerpb.GetNamespacesResponse, error) {
							return nil, io.EOF
						},
					}
				},
			},
			want: want{
				resp: &observerpb.GetNamespacesResponse{
					Namespaces: []*observerpb.Namespace{},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=<nil> peer=noip`,
				},
			},
		},
		{
			name: "2 connected peer, 1 unreachable peer",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
						{
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						},
						{
							Peer: peerTypes.Peer{
								Name: "three",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.3"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnGetNamespaces: func(_ context.Context, in *observerpb.GetNamespacesRequest, _ ...grpc.CallOption) (*observerpb.GetNamespacesResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.GetNamespacesResponse{
									Namespaces: []*observerpb.Namespace{
										{
											Namespace: "zzz",
											Cluster:   "some-cluster",
										},
										{
											Namespace: "aaa",
											Cluster:   "some-cluster",
										},
										{
											Namespace: "bbb",
											Cluster:   "some-cluster",
										},
									},
								}, nil
							case "three":
								return &observerpb.GetNamespacesResponse{
									Namespaces: []*observerpb.Namespace{
										{
											Namespace: "zzz",
											Cluster:   "some-cluster",
										},
										{
											Namespace: "ccc",
											Cluster:   "some-cluster",
										},
										{
											Namespace: "ddd",
											Cluster:   "some-cluster",
										},
									},
								}, nil
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.GetNamespacesResponse{
					Namespaces: []*observerpb.Namespace{
						{
							Namespace: "aaa",
							Cluster:   "some-cluster",
						},
						{
							Namespace: "bbb",
							Cluster:   "some-cluster",
						},
						{
							Namespace: "ccc",
							Cluster:   "some-cluster",
						},
						{
							Namespace: "ddd",
							Cluster:   "some-cluster",
						},
						{
							Namespace: "zzz",
							Cluster:   "some-cluster",
						},
					},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=192.0.2.2:4244 peer=two`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := slog.New(
				slog.NewTextHandler(&buf,
					&slog.HandlerOptions{
						ReplaceAttr: logging.ReplaceAttrFnWithoutTimestamp,
					},
				),
			)
			srv, err := NewServer(
				tt.plr,
				WithLogger(logger),
				withObserverClientBuilder(tt.ocb),
			)
			assert.NoError(t, err)
			got, err := srv.GetNamespaces(t.Context(), tt.req)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.resp, got)
			out := buf.String()
			for _, msg := range tt.want.log {
				assert.Contains(t, out, msg)
			}
		})
	}
}

func TestServerStatus(t *testing.T) {
	type want struct {
		resp *observerpb.ServerStatusResponse
		err  error
		log  []string
	}
	tests := []struct {
		name string
		plr  PeerLister
		ocb  observerClientBuilder
		req  *observerpb.ServerStatusRequest
		want want
	}{
		{
			name: "1 peer without address",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name:    "noip",
								Address: nil,
							},
							Conn: nil,
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{},
			want: want{
				resp: &observerpb.ServerStatusResponse{
					Version:             "hubble-relay",
					NumFlows:            0,
					MaxFlows:            0,
					SeenFlows:           0,
					UptimeNs:            0,
					FlowsRate:           0,
					NumConnectedNodes:   &wrapperspb.UInt32Value{Value: 0},
					NumUnavailableNodes: &wrapperspb.UInt32Value{Value: 1},
					UnavailableNodes:    []string{"noip"},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=<nil> peer=noip`,
				},
			},
		}, {
			name: "2 connected peers",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.ServerStatusResponse{
									NumFlows:  1111,
									MaxFlows:  1111,
									SeenFlows: 1111,
									FlowsRate: 1,
									UptimeNs:  111111111,
								}, nil
							case "two":
								return &observerpb.ServerStatusResponse{
									NumFlows:  2222,
									MaxFlows:  2222,
									SeenFlows: 2222,
									FlowsRate: 2,
									UptimeNs:  222222222,
								}, nil
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.ServerStatusResponse{
					Version:             "hubble-relay",
					NumFlows:            3333,
					MaxFlows:            3333,
					SeenFlows:           3333,
					FlowsRate:           3,
					UptimeNs:            222222222,
					NumConnectedNodes:   &wrapperspb.UInt32Value{Value: 2},
					NumUnavailableNodes: &wrapperspb.UInt32Value{Value: 0},
				},
			},
		}, {
			name: "1 connected peer, 1 unreachable peer",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.Ready
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							switch p.Name {
							case "one":
								return &observerpb.ServerStatusResponse{
									NumFlows:  1111,
									MaxFlows:  1111,
									SeenFlows: 1111,
									FlowsRate: 1,
									UptimeNs:  111111111,
								}, nil
							default:
								return nil, io.EOF
							}
						},
					}
				},
			},
			want: want{
				resp: &observerpb.ServerStatusResponse{
					Version:             "hubble-relay",
					NumFlows:            1111,
					MaxFlows:            1111,
					SeenFlows:           1111,
					FlowsRate:           1,
					UptimeNs:            111111111,
					NumConnectedNodes:   &wrapperspb.UInt32Value{Value: 1},
					NumUnavailableNodes: &wrapperspb.UInt32Value{Value: 1},
					UnavailableNodes:    []string{"two"},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=192.0.2.2:4244 peer=two`,
				},
			},
		}, {
			name: "2 unreachable peers",
			plr: &testutils.FakePeerLister{
				OnList: func() []poolTypes.Peer {
					return []poolTypes.Peer{
						{
							Peer: peerTypes.Peer{
								Name: "one",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.1"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						}, {
							Peer: peerTypes.Peer{
								Name: "two",
								Address: &net.TCPAddr{
									IP:   net.ParseIP("192.0.2.2"),
									Port: defaults.ServerPort,
								},
							},
							Conn: &testutils.FakeClientConn{
								OnGetState: func() connectivity.State {
									return connectivity.TransientFailure
								},
							},
						},
					}
				},
			},
			ocb: fakeObserverClientBuilder{
				onObserverClient: func(p *poolTypes.Peer) observerpb.ObserverClient {
					return &testutils.FakeObserverClient{
						OnServerStatus: func(_ context.Context, in *observerpb.ServerStatusRequest, _ ...grpc.CallOption) (*observerpb.ServerStatusResponse, error) {
							return nil, io.EOF
						},
					}
				},
			},
			want: want{
				resp: &observerpb.ServerStatusResponse{
					Version:             "hubble-relay",
					NumFlows:            0,
					MaxFlows:            0,
					SeenFlows:           0,
					UptimeNs:            0,
					NumConnectedNodes:   &wrapperspb.UInt32Value{Value: 0},
					NumUnavailableNodes: &wrapperspb.UInt32Value{Value: 2},
					UnavailableNodes:    []string{"one", "two"},
				},
				log: []string{
					`level=info msg="No connection to peer, skipping" address=192.0.2.1:4244 peer=one`,
					`level=info msg="No connection to peer, skipping" address=192.0.2.2:4244 peer=two`,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			logger := slog.New(
				slog.NewTextHandler(&buf,
					&slog.HandlerOptions{
						ReplaceAttr: logging.ReplaceAttrFnWithoutTimestamp,
					},
				),
			)
			srv, err := NewServer(
				tt.plr,
				WithLogger(logger),
				withObserverClientBuilder(tt.ocb),
			)
			assert.NoError(t, err)
			got, err := srv.ServerStatus(t.Context(), tt.req)
			assert.Equal(t, tt.want.err, err)
			assert.Equal(t, tt.want.resp, got)
			out := buf.String()
			for _, msg := range tt.want.log {
				assert.Contains(t, out, msg)
			}
		})
	}
}

type fakeObserverClientBuilder struct {
	onObserverClient func(*poolTypes.Peer) observerpb.ObserverClient
}

func (b fakeObserverClientBuilder) observerClient(p *poolTypes.Peer) observerpb.ObserverClient {
	if b.onObserverClient != nil {
		return b.onObserverClient(p)
	}
	panic("OnObserverClient not set")
}
