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

// FiltersSecurityGroup One or more filters.
type FiltersSecurityGroup struct {
	// The account IDs of the owners of the security groups.
	AccountIds []string `json:"AccountIds,omitempty"`
	// The descriptions of the security groups.
	Descriptions []string `json:"Descriptions,omitempty"`
	// The account IDs that have been granted permissions.
	InboundRuleAccountIds []string `json:"InboundRuleAccountIds,omitempty"`
	// The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
	InboundRuleFromPortRanges []int32 `json:"InboundRuleFromPortRanges,omitempty"`
	// The IP ranges that have been granted permissions, in CIDR notation (for example, 10.0.0.0/24).
	InboundRuleIpRanges []string `json:"InboundRuleIpRanges,omitempty"`
	// The IP protocols for the permissions (`tcp` \\| `udp` \\| `icmp`, or a protocol number, or `-1` for all protocols).
	InboundRuleProtocols []string `json:"InboundRuleProtocols,omitempty"`
	// The IDs of the security groups that have been granted permissions.
	InboundRuleSecurityGroupIds []string `json:"InboundRuleSecurityGroupIds,omitempty"`
	// The names of the security groups that have been granted permissions.
	InboundRuleSecurityGroupNames []string `json:"InboundRuleSecurityGroupNames,omitempty"`
	// The ends of the port ranges for the TCP and UDP protocols, or the ICMP codes.
	InboundRuleToPortRanges []int32 `json:"InboundRuleToPortRanges,omitempty"`
	// The IDs of the Nets specified when the security groups were created.
	NetIds []string `json:"NetIds,omitempty"`
	// The account IDs that have been granted permissions.
	OutboundRuleAccountIds []string `json:"OutboundRuleAccountIds,omitempty"`
	// The beginnings of the port ranges for the TCP and UDP protocols, or the ICMP type numbers.
	OutboundRuleFromPortRanges []int32 `json:"OutboundRuleFromPortRanges,omitempty"`
	// The IP ranges that have been granted permissions, in CIDR notation (for example, 10.0.0.0/24).
	OutboundRuleIpRanges []string `json:"OutboundRuleIpRanges,omitempty"`
	// The IP protocols for the permissions (`tcp` \\| `udp` \\| `icmp`, or a protocol number, or `-1` for all protocols).
	OutboundRuleProtocols []string `json:"OutboundRuleProtocols,omitempty"`
	// The IDs of the security groups that have been granted permissions.
	OutboundRuleSecurityGroupIds []string `json:"OutboundRuleSecurityGroupIds,omitempty"`
	// The names of the security groups that have been granted permissions.
	OutboundRuleSecurityGroupNames []string `json:"OutboundRuleSecurityGroupNames,omitempty"`
	// The ends of the port ranges for the TCP and UDP protocols, or the ICMP codes.
	OutboundRuleToPortRanges []int32 `json:"OutboundRuleToPortRanges,omitempty"`
	// The IDs of the security groups.
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`
	// The names of the security groups.
	SecurityGroupNames []string `json:"SecurityGroupNames,omitempty"`
	// The keys of the tags associated with the security groups.
	TagKeys []string `json:"TagKeys,omitempty"`
	// The values of the tags associated with the security groups.
	TagValues []string `json:"TagValues,omitempty"`
	// The key/value combination of the tags associated with the security groups, in the following format: &quot;Filters&quot;:{&quot;Tags&quot;:[&quot;TAGKEY=TAGVALUE&quot;]}.
	Tags []string `json:"Tags,omitempty"`
}
