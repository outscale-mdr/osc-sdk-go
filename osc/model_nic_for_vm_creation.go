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

// NicForVmCreation Information about the network interface card (NIC) when creating a virtual machine (VM).
type NicForVmCreation struct {
	// By default or if set to true, the NIC is deleted when the VM is terminated. You can specify this parameter only for a new NIC. To modify this value for an existing NIC, see [UpdateNic](#updatenic).
	DeleteOnVmDeletion bool `json:"DeleteOnVmDeletion,omitempty"`
	// The description of the NIC, if you are creating a NIC when creating the VM.
	Description string `json:"Description,omitempty"`
	// The index of the VM device for the NIC attachment (between 0 and 7, both included). This parameter is required if you create a NIC when creating the VM.
	DeviceNumber int32 `json:"DeviceNumber,omitempty"`
	// The ID of the NIC, if you are attaching an existing NIC when creating a VM.
	NicId string `json:"NicId,omitempty"`
	// One or more private IP addresses to assign to the NIC, if you create a NIC when creating a VM. Only one private IP address can be the primary private IP address.
	PrivateIps []PrivateIpLight `json:"PrivateIps,omitempty"`
	// The number of secondary private IP addresses, if you create a NIC when creating a VM. This parameter cannot be specified if you specified more than one private IP address in the `PrivateIps` parameter.
	SecondaryPrivateIpCount int32 `json:"SecondaryPrivateIpCount,omitempty"`
	// One or more IDs of security groups for the NIC, if you acreate a NIC when creating a VM.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// The ID of the Subnet for the NIC, if you create a NIC when creating a VM.
	SubnetId string `json:"SubnetId,omitempty"`
}
