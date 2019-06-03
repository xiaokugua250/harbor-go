/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type Resource struct {
	// The replication policy list.
	ReplicationPolicies []ReplicationPolicy `json:"replication_policies,omitempty"`
}
