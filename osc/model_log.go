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

// Log Information about the log.
type Log struct {
	// The account ID of the logged call.
	AccountId string `json:"AccountId,omitempty"`
	// The duration of the logged call, in microseconds.
	CallDuration int32 `json:"CallDuration,omitempty"`
	// The access key used for the logged call.
	QueryAccessKey string `json:"QueryAccessKey,omitempty"`
	// The name of the API used by the logged call (always `oapi` for the OUTSCALE API).
	QueryApiName string `json:"QueryApiName,omitempty"`
	// The version of the API used by the logged call.
	QueryApiVersion string `json:"QueryApiVersion,omitempty"`
	// The name of the logged call.
	QueryCallName string `json:"QueryCallName,omitempty"`
	// The date of the logged call, in ISO 8601 format.
	QueryDate string `json:"QueryDate,omitempty"`
	// The raw header of the HTTP request of the logged call.
	QueryHeaderRaw string `json:"QueryHeaderRaw,omitempty"`
	// The size of the raw header of the HTTP request of the logged call, in bytes.
	QueryHeaderSize int32 `json:"QueryHeaderSize,omitempty"`
	// The IP address used for the logged call.
	QueryIpAddress string `json:"QueryIpAddress,omitempty"`
	// The raw payload of the HTTP request of the logged call.
	QueryPayloadRaw string `json:"QueryPayloadRaw,omitempty"`
	// The size of the raw payload of the HTTP request of the logged call, in bytes.
	QueryPayloadSize int32 `json:"QueryPayloadSize,omitempty"`
	// The user agent of the HTTP request of the logged call.
	QueryUserAgent string `json:"QueryUserAgent,omitempty"`
	// The request ID provided in the response of the logged call.
	RequestId string `json:"RequestId,omitempty"`
	// The size of the response of the logged call, in bytes.
	ResponseSize int32 `json:"ResponseSize,omitempty"`
	// The HTTP status code of the response of the logged call.
	ResponseStatusCode int32 `json:"ResponseStatusCode,omitempty"`
}
