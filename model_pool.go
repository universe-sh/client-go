/*
 * Universe.sh houston API
 *
 * Universe.sh houston API
 *
 * API version: 1.0.0
 * Contact: oss@universe.sh
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Pool struct for Pool
type Pool struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	LaunchSpecification PoolLaunchSpecification `json:"launch_specification,omitempty"`
	Autoscaling PoolAutoscaling `json:"autoscaling,omitempty"`
	Actived bool `json:"actived,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
}
