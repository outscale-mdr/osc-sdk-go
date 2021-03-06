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

// FiltersVmType One or more filters.
type FiltersVmType struct {
	// Indicates whether the VM is optimized for BSU I/O.
	BsuOptimized bool `json:"BsuOptimized,omitempty"`
	// The amounts of memory, in gibibytes (GiB).
	MemorySizes []float32 `json:"MemorySizes,omitempty"`
	// The numbers of vCores.
	VcoreCounts []int32 `json:"VcoreCounts,omitempty"`
	// The names of the VM types. For more information, see [Instance Types](https://wiki.outscale.net/display/EN/Instance+Types).
	VmTypeNames []string `json:"VmTypeNames,omitempty"`
	// The maximum number of ephemeral storage disks.
	VolumeCounts []int32 `json:"VolumeCounts,omitempty"`
	// The size of one ephemeral storage disk, in gibibytes (GiB).
	VolumeSizes []int32 `json:"VolumeSizes,omitempty"`
}
