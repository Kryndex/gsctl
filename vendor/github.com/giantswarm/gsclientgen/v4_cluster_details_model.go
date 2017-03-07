/*
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * OpenAPI spec version: legacy
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package gsclientgen

type V4ClusterDetailsModel struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ApiEndpoint string `json:"api_endpoint,omitempty"`

	CreateDate string `json:"create_date,omitempty"`

	Owner string `json:"owner,omitempty"`

	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	Workers []V4NodeDefinitionModel `json:"workers,omitempty"`

	Masters []V4NodeDefinitionModel `json:"masters,omitempty"`
}