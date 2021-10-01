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

// DirectLink Information about the DirectLink.
type DirectLink struct {
	// The account ID of the owner of the DirectLink.
	AccountId string `json:"AccountId,omitempty"`
	// The physical link bandwidth (either 1 Gbps or 10 Gbps).
	Bandwidth string `json:"Bandwidth,omitempty"`
	// The ID of the DirectLink (for example, dxcon-xxxxxxxx).
	DirectLinkId string `json:"DirectLinkId,omitempty"`
	// The name of the DirectLink.
	DirectLinkName string `json:"DirectLinkName,omitempty"`
	// The datacenter where the DirectLink is located.
	Location string `json:"Location,omitempty"`
	// The Region in which the DirectLink has been created.
	RegionName string `json:"RegionName,omitempty"`
	// The state of the DirectLink.<br /> * `requested`: The DirectLink is requested but the request has not been validated yet.<br /> * `pending`: The DirectLink request has been validated. It remains in the `pending` state until you establish the physical link.<br /> * `available`: The physical link is established and the connection is ready to use.<br /> * `deleting`: The deletion process is in progress.<br /> * `deleted`: The DirectLink is deleted.
	State string `json:"State,omitempty"`
}
