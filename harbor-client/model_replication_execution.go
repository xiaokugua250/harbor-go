/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.7.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

// The replication execution
type ReplicationExecution struct {
	// The ID
	Id int32 `json:"id,omitempty"`
	// The policy ID
	PolicyId int32 `json:"policy_id,omitempty"`
	// The status
	Status string `json:"status,omitempty"`
	// The status text
	StatusText string `json:"status_text,omitempty"`
	// The trigger mode
	Trigger string `json:"trigger,omitempty"`
	// The total count of all tasks
	Total int32 `json:"total,omitempty"`
	// The count of failed tasks
	Failed int32 `json:"failed,omitempty"`
	// The count of succeed tasks
	Succeed int32 `json:"succeed,omitempty"`
	// The count of in_progress tasks
	InProgress int32 `json:"in_progress,omitempty"`
	// The count of stopped tasks
	Stopped int32 `json:"stopped,omitempty"`
	// The start time
	StartTime string `json:"start_time,omitempty"`
	// The end time
	EndTime string `json:"end_time,omitempty"`
}
