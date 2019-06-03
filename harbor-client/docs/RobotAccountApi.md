# \RobotAccountApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectsProjectIdRobotsGet**](RobotAccountApi.md#ProjectsProjectIdRobotsGet) | **Get** /projects/{project_id}/robots | Get all robot accounts of specified project
[**ProjectsProjectIdRobotsPost**](RobotAccountApi.md#ProjectsProjectIdRobotsPost) | **Post** /projects/{project_id}/robots | Create a robot account for project
[**ProjectsProjectIdRobotsRobotIdDelete**](RobotAccountApi.md#ProjectsProjectIdRobotsRobotIdDelete) | **Delete** /projects/{project_id}/robots/{robot_id} | Delete the specified robot account
[**ProjectsProjectIdRobotsRobotIdGet**](RobotAccountApi.md#ProjectsProjectIdRobotsRobotIdGet) | **Get** /projects/{project_id}/robots/{robot_id} | Return the infor of the specified robot account.
[**ProjectsProjectIdRobotsRobotIdPut**](RobotAccountApi.md#ProjectsProjectIdRobotsRobotIdPut) | **Put** /projects/{project_id}/robots/{robot_id} | Update status of robot account.


# **ProjectsProjectIdRobotsGet**
> []RobotAccount ProjectsProjectIdRobotsGet(ctx, projectId)
Get all robot accounts of specified project

Get all robot accounts of specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 

### Return type

[**[]RobotAccount**](RobotAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsPost**
> RobotAccountPostRep ProjectsProjectIdRobotsPost(ctx, projectId, robot)
Create a robot account for project

Create a robot account for project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **robot** | [**RobotAccountCreate**](RobotAccountCreate.md)| Request body of creating a robot account. | 

### Return type

[**RobotAccountPostRep**](RobotAccountPostRep.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsRobotIdDelete**
> ProjectsProjectIdRobotsRobotIdDelete(ctx, projectId, robotId)
Delete the specified robot account

Delete the specified robot account

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **robotId** | **int64**| The ID of robot account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsRobotIdGet**
> RobotAccount ProjectsProjectIdRobotsRobotIdGet(ctx, projectId, robotId)
Return the infor of the specified robot account.

Return the infor of the specified robot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **robotId** | **int64**| The ID of robot account. | 

### Return type

[**RobotAccount**](RobotAccount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdRobotsRobotIdPut**
> ProjectsProjectIdRobotsRobotIdPut(ctx, projectId, robotId, robot)
Update status of robot account.

Used to disable/enable a specified robot account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **robotId** | **int64**| The ID of robot account. | 
  **robot** | [**RobotAccountUpdate**](RobotAccountUpdate.md)| Request body of enable/disable a robot account. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

