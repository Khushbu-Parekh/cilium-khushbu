// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package builder

import (
	_ "embed"

	"github.com/cilium/cilium/cilium-cli/connectivity/check"
	"github.com/cilium/cilium/cilium-cli/connectivity/tests"
	"github.com/cilium/cilium/cilium-cli/utils/features"
)

//go:embed manifests/client-egress-l7-http-method.yaml
var clientEgressL7HTTPMethodPolicyYAML string

//go:embed manifests/client-egress-l7-http-method-port-range.yaml
var clientEgressL7HTTPMethodPolicyPortRangeYAML string

type clientEgressL7Method struct{}

func (t clientEgressL7Method) build(ct *check.ConnectivityTest, templates map[string]string) {
	clientEgressL7MethodTest(ct, templates, false)
	if ct.Features[features.L7PortRanges].Enabled {
		clientEgressL7MethodTest(ct, templates, true)
	}
}

func clientEgressL7MethodTest(ct *check.ConnectivityTest, templates map[string]string, portRanges bool) {
	testName := "client-egress-l7-method"
	yamlFile := clientEgressL7HTTPMethodPolicyYAML
	if portRanges {
		testName = "client-egress-l7-method-port-range"
		yamlFile = clientEgressL7HTTPMethodPolicyPortRangeYAML
	}
	// Test L7 HTTP with different methods introspection using an egress policy on the clients.
	newTest(testName, ct).
		WithFeatureRequirements(features.RequireEnabled(features.L7Proxy)).
		WithCiliumPolicy(templates["clientEgressOnlyDNSPolicyYAML"]). // DNS resolution only
		WithCiliumPolicy(yamlFile).                                   // L7 allow policy with HTTP introspection (POST only)
		WithScenarios(
			tests.PodToPodWithEndpoints(tests.WithMethod("POST"), tests.WithDestinationLabelsOption(map[string]string{"other": "echo"})),
			tests.PodToPodWithEndpoints(tests.WithDestinationLabelsOption(map[string]string{"first": "echo"})),
		).
		WithExpectations(func(a *check.Action) (egress, ingress check.Result) {
			if a.Source().HasLabel("other", "client") && // Only client2 is allowed to make HTTP calls.
				(a.Destination().Port() == 8080) { // port 8080 is traffic to echo Pod.
				if a.Destination().HasLabel("other", "echo") { //we are POSTing only other echo
					egress = check.ResultOK

					egress.HTTP = check.HTTP{
						Method: "POST",
					}
					return egress, check.ResultNone
				}
				// Else expect HTTP drop by proxy
				return check.ResultDropCurlHTTPError, check.ResultNone
			}
			return check.ResultDefaultDenyEgressDrop, check.ResultNone
		})
}
