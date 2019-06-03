/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

// The permission
type Permission struct {
	// The permission resoruce
	Resource string `json:"resource,omitempty"`
	// The permission action
	Action string `json:"action,omitempty"`
}
