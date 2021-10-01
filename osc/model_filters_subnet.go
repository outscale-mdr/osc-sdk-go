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

// FiltersSubnet One or more filters.
type FiltersSubnet struct {
	// The number of available IPs.
	AvailableIpsCounts []int32 `json:"AvailableIpsCounts,omitempty"`
	// The IP ranges in the Subnets, in CIDR notation (for example, 10.0.0.0/16).
	IpRanges []string `json:"IpRanges,omitempty"`
	// The IDs of the Nets in which the Subnets are.
	NetIds []string `json:"NetIds,omitempty"`
	// The states of the Subnets (`pending` \\| `available`).
	States []string `json:"States,omitempty"`
	// The IDs of the Subnets.
	SubnetIds []string `json:"SubnetIds,omitempty"`
	// The names of the Subregions in which the Subnets are located.
	SubregionNames []string `json:"SubregionNames,omitempty"`
	// The keys of the tags associated with the Subnets.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the Subnets.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the Subnets, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}