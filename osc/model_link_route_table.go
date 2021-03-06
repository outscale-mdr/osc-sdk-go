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

// LinkRouteTable One or more associations between the route table and the Subnets.
type LinkRouteTable struct {
	// The ID of the association between the route table and the Subnet.
	LinkRouteTableId string `json:"LinkRouteTableId,omitempty"`
	// If true, the route table is the main one.
	Main bool `json:"Main,omitempty"`
	// The ID of the route table.
	RouteTableId string `json:"RouteTableId,omitempty"`
	// The ID of the Subnet.
	SubnetId string `json:"SubnetId,omitempty"`
}
