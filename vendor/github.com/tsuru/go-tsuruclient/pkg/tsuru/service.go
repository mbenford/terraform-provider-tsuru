/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type Service struct {
	Service string `json:"service,omitempty"`

	// [deprecated]
	Instances []string `json:"instances,omitempty"`

	Plans []string `json:"plans,omitempty"`

	ServiceInstances []ServiceInstance `json:"service_instances,omitempty"`
}
