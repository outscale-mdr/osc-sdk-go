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

// PublicIp Information about the public IP.
type PublicIp struct {
	// (Required in a Net) The ID representing the association of the EIP with the VM or the NIC.
	LinkPublicIpId string `json:"LinkPublicIpId,omitempty"`
	// The account ID of the owner of the NIC.
	NicAccountId string `json:"NicAccountId,omitempty"`
	// The ID of the NIC the EIP is associated with (if any).
	NicId string `json:"NicId,omitempty"`
	// The private IP address associated with the EIP.
	PrivateIp string `json:"PrivateIp,omitempty"`
	// The External IP address (EIP) associated with the NAT service.
	PublicIp string `json:"PublicIp,omitempty"`
	// The allocation ID of the EIP associated with the NAT service.
	PublicIpId string `json:"PublicIpId,omitempty"`
	// One or more tags associated with the EIP.
	Tags []ResourceTag `json:"Tags,omitempty"`
	// The ID of the VM the External IP (EIP) is associated with (if any).
	VmId string `json:"VmId,omitempty"`
}
