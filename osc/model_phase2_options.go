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

// Phase2Options Information about Phase 2 of the Internet Key Exchange (IKE) negotiation.
type Phase2Options struct {
	// The Diffie-Hellman (DH) group numbers allowed for the VPN tunnel for phase 2.
	Phase2DhGroupNumbers []int32 `json:"Phase2DhGroupNumbers,omitempty"`
	// The encryption algorithms allowed for the VPN tunnel for phase 2.
	Phase2EncryptionAlgorithms []string `json:"Phase2EncryptionAlgorithms,omitempty"`
	// The integrity algorithms allowed for the VPN tunnel for phase 2.
	Phase2IntegrityAlgorithms []string `json:"Phase2IntegrityAlgorithms,omitempty"`
	// The lifetime for phase 2 of the Internet Key Exchange (IKE) negociation process, in seconds.
	Phase2LifetimeSeconds int32 `json:"Phase2LifetimeSeconds,omitempty"`
	// The pre-shared key to establish the initial authentication between the client gateway and the virtual gateway. This key can contain any character except line breaks and double quotes (&quot;).
	PreSharedKey string `json:"PreSharedKey,omitempty"`
}
