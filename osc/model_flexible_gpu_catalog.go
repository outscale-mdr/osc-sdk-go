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

// FlexibleGpuCatalog Information about the flexible GPU (fGPU) that is available in the public catalog.
type FlexibleGpuCatalog struct {
	// The generations of VMs that the fGPU is compatible with.
	Generations []string `json:"Generations,omitempty"`
	// The maximum number of VM vCores that the fGPU is compatible with.
	MaxCpu int32 `json:"MaxCpu,omitempty"`
	// The maximum amount of VM memory that the fGPU is compatible with.
	MaxRam int32 `json:"MaxRam,omitempty"`
	// The model of fGPU.
	ModelName string `json:"ModelName,omitempty"`
	// The amount of video RAM (VRAM) of the fGPU.
	VRam int32 `json:"VRam,omitempty"`
}
