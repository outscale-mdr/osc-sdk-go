/*
 * 3DS OUTSCALE API
 *
 * Welcome to the OUTSCALE API documentation.<br /><br />  The OUTSCALE API enables you to manage your resources in the OUTSCALE Cloud. This documentation describes the different actions available along with code examples.<br /><br />  Note that the OUTSCALE Cloud is compatible with Amazon Web Services (AWS) APIs, but some resources have different names in AWS than in the OUTSCALE API. You can find a list of the differences [here](https://wiki.outscale.net/display/EN/3DS+OUTSCALE+APIs+Reference).<br /><br />  You can also manage your resources using the [Cockpit](https://wiki.outscale.net/display/EN/About+Cockpit) web interface.
 *
 * API version: 1.15
 * Contact: support@outscale.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package osc

// ListenerForCreation Information about the listener to create.
type ListenerForCreation struct {
	// The port on which the back-end VM is listening (between `1` and `65535`, both included).
	BackendPort int32 `json:"BackendPort"`
	// The protocol for routing traffic to back-end VMs (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	BackendProtocol string `json:"BackendProtocol,omitempty"`
	// The port on which the load balancer is listening (between `1` and `65535`, both included).
	LoadBalancerPort int32 `json:"LoadBalancerPort"`
	// The routing protocol (`HTTP` \\| `HTTPS` \\| `TCP` \\| `SSL`).
	LoadBalancerProtocol string `json:"LoadBalancerProtocol"`
	// The OUTSCALE Resource Name (ORN) of the server certificate. For more information, see [Resource Identifiers > OUTSCALE Resource Names (ORNs)](https://wiki.outscale.net/display/EN/Resource+Identifiers#ResourceIdentifiers-ORNFormat).
	ServerCertificateId string `json:"ServerCertificateId,omitempty"`
}
