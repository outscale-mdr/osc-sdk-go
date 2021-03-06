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

// CreateLoadBalancerPolicyRequest struct for CreateLoadBalancerPolicyRequest
type CreateLoadBalancerPolicyRequest struct {
	// The lifetime of the cookie, in seconds. If not specified, the default value of this parameter is 1, which means that the sticky session lasts for the duration of the browser session.
	CookieExpirationPeriod int32 `json:"CookieExpirationPeriod,omitempty"`
	// The name of the application cookie used for stickiness. This parameter is required if you create a stickiness policy based on an application-generated cookie.
	CookieName string `json:"CookieName,omitempty"`
	// If true, checks whether you have the required permissions to perform the action.
	DryRun bool `json:"DryRun,omitempty"`
	// The name of the load balancer for which you want to create a policy.
	LoadBalancerName string `json:"LoadBalancerName"`
	// The name of the policy. This name must be unique and consist of alphanumeric characters and dashes (-).
	PolicyName string `json:"PolicyName"`
	// The type of stickiness policy you want to create: `app` or `load_balancer`.
	PolicyType string `json:"PolicyType"`
}
