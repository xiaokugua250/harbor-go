/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type HasAdminRole struct {
	// 1-has admin, 0-not.
	HasAdminRole bool `json:"has_admin_role,omitempty"`
}
