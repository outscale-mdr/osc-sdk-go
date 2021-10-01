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

// AccessLog Information about access logs.
type AccessLog struct {
	// If true, access logs are enabled for your load balancer. If false, they are not. If you set this to true in your request, the `OsuBucketName` parameter is required.
	IsEnabled bool `json:"IsEnabled,omitempty"`
	// The name of the OOS bucket for the access logs.
	OsuBucketName string `json:"OsuBucketName,omitempty"`
	// The path to the folder of the access logs in your OOS bucket (by default, the `root` level of your bucket).
	OsuBucketPrefix string `json:"OsuBucketPrefix,omitempty"`
	// The time interval for the publication of access logs in the OOS bucket, in minutes. This value can be either 5 or 60 (by default, 60).
	PublicationInterval int32 `json:"PublicationInterval,omitempty"`
}
