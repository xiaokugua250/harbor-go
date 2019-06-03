# \ProductsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChartrepoChartsPost**](ProductsApi.md#ChartrepoChartsPost) | **Post** /chartrepo/charts | Upload a chart file to the defult &#39;library&#39; project.
[**ChartrepoHealthGet**](ProductsApi.md#ChartrepoHealthGet) | **Get** /chartrepo/health | Check the health of chart repository service.
[**ChartrepoRepoChartsGet**](ProductsApi.md#ChartrepoRepoChartsGet) | **Get** /chartrepo/{repo}/charts | Get all the charts under the specified project
[**ChartrepoRepoChartsNameDelete**](ProductsApi.md#ChartrepoRepoChartsNameDelete) | **Delete** /chartrepo/{repo}/charts/{name} | Delete all the versions of the specified chart
[**ChartrepoRepoChartsNameGet**](ProductsApi.md#ChartrepoRepoChartsNameGet) | **Get** /chartrepo/{repo}/charts/{name} | Get all the versions of the specified chart
[**ChartrepoRepoChartsNameVersionDelete**](ProductsApi.md#ChartrepoRepoChartsNameVersionDelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version} | Delete the specified chart version
[**ChartrepoRepoChartsNameVersionGet**](ProductsApi.md#ChartrepoRepoChartsNameVersionGet) | **Get** /chartrepo/{repo}/charts/{name}/{version} | Get the specified chart version
[**ChartrepoRepoChartsNameVersionLabelsGet**](ProductsApi.md#ChartrepoRepoChartsNameVersionLabelsGet) | **Get** /chartrepo/{repo}/charts/{name}/{version}/labels | Return the attahced labels of chart.
[**ChartrepoRepoChartsNameVersionLabelsIdDelete**](ProductsApi.md#ChartrepoRepoChartsNameVersionLabelsIdDelete) | **Delete** /chartrepo/{repo}/charts/{name}/{version}/labels/{id} | Remove label from chart.
[**ChartrepoRepoChartsNameVersionLabelsPost**](ProductsApi.md#ChartrepoRepoChartsNameVersionLabelsPost) | **Post** /chartrepo/{repo}/charts/{name}/{version}/labels | Mark label to chart.
[**ChartrepoRepoChartsPost**](ProductsApi.md#ChartrepoRepoChartsPost) | **Post** /chartrepo/{repo}/charts | Upload a chart file to the specified project.
[**ChartrepoRepoProvPost**](ProductsApi.md#ChartrepoRepoProvPost) | **Post** /chartrepo/{repo}/prov | Upload a provance file to the specified project.
[**ConfigurationsGet**](ProductsApi.md#ConfigurationsGet) | **Get** /configurations | Get system configurations.
[**ConfigurationsPut**](ProductsApi.md#ConfigurationsPut) | **Put** /configurations | Modify system configurations.
[**EmailPingPost**](ProductsApi.md#EmailPingPost) | **Post** /email/ping | Test connection and authentication with email server.
[**HealthGet**](ProductsApi.md#HealthGet) | **Get** /health | Health check API
[**InternalSyncregistryPost**](ProductsApi.md#InternalSyncregistryPost) | **Post** /internal/syncregistry | Sync repositories from registry to DB.
[**LabelsGet**](ProductsApi.md#LabelsGet) | **Get** /labels | List labels according to the query strings.
[**LabelsIdDelete**](ProductsApi.md#LabelsIdDelete) | **Delete** /labels/{id} | Delete the label specified by ID.
[**LabelsIdGet**](ProductsApi.md#LabelsIdGet) | **Get** /labels/{id} | Get the label specified by ID.
[**LabelsIdPut**](ProductsApi.md#LabelsIdPut) | **Put** /labels/{id} | Update the label properties.
[**LabelsIdResourcesGet**](ProductsApi.md#LabelsIdResourcesGet) | **Get** /labels/{id}/resources | Get the resources that the label is referenced by.
[**LabelsPost**](ProductsApi.md#LabelsPost) | **Post** /labels | Post creates a label
[**LdapGroupsSearchGet**](ProductsApi.md#LdapGroupsSearchGet) | **Get** /ldap/groups/search | Search available ldap groups.
[**LdapPingPost**](ProductsApi.md#LdapPingPost) | **Post** /ldap/ping | Ping available ldap service.
[**LdapUsersImportPost**](ProductsApi.md#LdapUsersImportPost) | **Post** /ldap/users/import | Import selected available ldap users.
[**LdapUsersSearchGet**](ProductsApi.md#LdapUsersSearchGet) | **Get** /ldap/users/search | Search available ldap users.
[**LogsGet**](ProductsApi.md#LogsGet) | **Get** /logs | Get recent logs of the projects which the user is a member of
[**ProjectsGet**](ProductsApi.md#ProjectsGet) | **Get** /projects | List projects
[**ProjectsHead**](ProductsApi.md#ProjectsHead) | **Head** /projects | Check if the project name user provided already exists.
[**ProjectsPost**](ProductsApi.md#ProjectsPost) | **Post** /projects | Create a new project.
[**ProjectsProjectIdDelete**](ProductsApi.md#ProjectsProjectIdDelete) | **Delete** /projects/{project_id} | Delete project by projectID
[**ProjectsProjectIdGet**](ProductsApi.md#ProjectsProjectIdGet) | **Get** /projects/{project_id} | Return specific project detail information
[**ProjectsProjectIdLogsGet**](ProductsApi.md#ProjectsProjectIdLogsGet) | **Get** /projects/{project_id}/logs | Get access logs accompany with a relevant project.
[**ProjectsProjectIdMembersGet**](ProductsApi.md#ProjectsProjectIdMembersGet) | **Get** /projects/{project_id}/members | Get all project member information
[**ProjectsProjectIdMembersMidDelete**](ProductsApi.md#ProjectsProjectIdMembersMidDelete) | **Delete** /projects/{project_id}/members/{mid} | Delete project member
[**ProjectsProjectIdMembersMidGet**](ProductsApi.md#ProjectsProjectIdMembersMidGet) | **Get** /projects/{project_id}/members/{mid} | Get the project member information
[**ProjectsProjectIdMembersMidPut**](ProductsApi.md#ProjectsProjectIdMembersMidPut) | **Put** /projects/{project_id}/members/{mid} | Update project member
[**ProjectsProjectIdMembersPost**](ProductsApi.md#ProjectsProjectIdMembersPost) | **Post** /projects/{project_id}/members | Create project member
[**ProjectsProjectIdMetadatasGet**](ProductsApi.md#ProjectsProjectIdMetadatasGet) | **Get** /projects/{project_id}/metadatas | Get project metadata.
[**ProjectsProjectIdMetadatasMetaNameDelete**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNameDelete) | **Delete** /projects/{project_id}/metadatas/{meta_name} | Delete metadata of a project
[**ProjectsProjectIdMetadatasMetaNameGet**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNameGet) | **Get** /projects/{project_id}/metadatas/{meta_name} | Get project metadata
[**ProjectsProjectIdMetadatasMetaNamePut**](ProductsApi.md#ProjectsProjectIdMetadatasMetaNamePut) | **Put** /projects/{project_id}/metadatas/{meta_name} | Update metadata of a project.
[**ProjectsProjectIdMetadatasPost**](ProductsApi.md#ProjectsProjectIdMetadatasPost) | **Post** /projects/{project_id}/metadatas | Add metadata for the project.
[**ProjectsProjectIdPut**](ProductsApi.md#ProjectsProjectIdPut) | **Put** /projects/{project_id} | Update properties for a selected project.
[**ProjectsProjectIdRobotsGet**](ProductsApi.md#ProjectsProjectIdRobotsGet) | **Get** /projects/{project_id}/robots | Get all robot accounts of specified project
[**ProjectsProjectIdRobotsPost**](ProductsApi.md#ProjectsProjectIdRobotsPost) | **Post** /projects/{project_id}/robots | Create a robot account for project
[**ProjectsProjectIdRobotsRobotIdDelete**](ProductsApi.md#ProjectsProjectIdRobotsRobotIdDelete) | **Delete** /projects/{project_id}/robots/{robot_id} | Delete the specified robot account
[**ProjectsProjectIdRobotsRobotIdGet**](ProductsApi.md#ProjectsProjectIdRobotsRobotIdGet) | **Get** /projects/{project_id}/robots/{robot_id} | Return the infor of the specified robot account.
[**ProjectsProjectIdRobotsRobotIdPut**](ProductsApi.md#ProjectsProjectIdRobotsRobotIdPut) | **Put** /projects/{project_id}/robots/{robot_id} | Update status of robot account.
[**RegistriesGet**](ProductsApi.md#RegistriesGet) | **Get** /registries | List registries.
[**RegistriesIdDelete**](ProductsApi.md#RegistriesIdDelete) | **Delete** /registries/{id} | Delete specific registry.
[**RegistriesIdGet**](ProductsApi.md#RegistriesIdGet) | **Get** /registries/{id} | Get registry.
[**RegistriesIdInfoGet**](ProductsApi.md#RegistriesIdInfoGet) | **Get** /registries/{id}/info | Get registry info.
[**RegistriesIdNamespaceGet**](ProductsApi.md#RegistriesIdNamespaceGet) | **Get** /registries/{id}/namespace | List namespaces of registry
[**RegistriesIdPut**](ProductsApi.md#RegistriesIdPut) | **Put** /registries/{id} | Update a given registry.
[**RegistriesPingPost**](ProductsApi.md#RegistriesPingPost) | **Post** /registries/ping | Ping status of a registry.
[**RegistriesPost**](ProductsApi.md#RegistriesPost) | **Post** /registries | Create a new registry.
[**ReplicationAdaptersGet**](ProductsApi.md#ReplicationAdaptersGet) | **Get** /replication/adapters | List supported adapters.
[**ReplicationExecutionsGet**](ProductsApi.md#ReplicationExecutionsGet) | **Get** /replication/executions | List replication executions.
[**ReplicationExecutionsIdGet**](ProductsApi.md#ReplicationExecutionsIdGet) | **Get** /replication/executions/{id} | Get the execution of the replication.
[**ReplicationExecutionsIdPut**](ProductsApi.md#ReplicationExecutionsIdPut) | **Put** /replication/executions/{id} | Stop the execution of the replication.
[**ReplicationExecutionsIdTasksGet**](ProductsApi.md#ReplicationExecutionsIdTasksGet) | **Get** /replication/executions/{id}/tasks | Get the task list of one execution.
[**ReplicationExecutionsIdTasksTaskIdLogGet**](ProductsApi.md#ReplicationExecutionsIdTasksTaskIdLogGet) | **Get** /replication/executions/{id}/tasks/{task_id}/log | Get the log of one task.
[**ReplicationExecutionsPost**](ProductsApi.md#ReplicationExecutionsPost) | **Post** /replication/executions | Start one execution of the replication.
[**ReplicationPoliciesGet**](ProductsApi.md#ReplicationPoliciesGet) | **Get** /replication/policies | List replication policies
[**ReplicationPoliciesIdDelete**](ProductsApi.md#ReplicationPoliciesIdDelete) | **Delete** /replication/policies/{id} | Delete the replication policy specified by ID.
[**ReplicationPoliciesIdGet**](ProductsApi.md#ReplicationPoliciesIdGet) | **Get** /replication/policies/{id} | Get replication policy.
[**ReplicationPoliciesIdPut**](ProductsApi.md#ReplicationPoliciesIdPut) | **Put** /replication/policies/{id} | Update the replication policy
[**ReplicationPoliciesPost**](ProductsApi.md#ReplicationPoliciesPost) | **Post** /replication/policies | Create a replication policy
[**RepositoriesGet**](ProductsApi.md#RepositoriesGet) | **Get** /repositories | Get repositories accompany with relevant project and repo name.
[**RepositoriesRepoNameDelete**](ProductsApi.md#RepositoriesRepoNameDelete) | **Delete** /repositories/{repo_name} | Delete a repository.
[**RepositoriesRepoNameLabelsGet**](ProductsApi.md#RepositoriesRepoNameLabelsGet) | **Get** /repositories/{repo_name}/labels | Get labels of a repository.
[**RepositoriesRepoNameLabelsLabelIdDelete**](ProductsApi.md#RepositoriesRepoNameLabelsLabelIdDelete) | **Delete** /repositories/{repo_name}/labels/{label_id} | Delete label from the repository.
[**RepositoriesRepoNameLabelsPost**](ProductsApi.md#RepositoriesRepoNameLabelsPost) | **Post** /repositories/{repo_name}/labels | Add a label to the repository.
[**RepositoriesRepoNamePut**](ProductsApi.md#RepositoriesRepoNamePut) | **Put** /repositories/{repo_name} | Update description of the repository.
[**RepositoriesRepoNameSignaturesGet**](ProductsApi.md#RepositoriesRepoNameSignaturesGet) | **Get** /repositories/{repo_name}/signatures | Get signature information of a repository
[**RepositoriesRepoNameTagsGet**](ProductsApi.md#RepositoriesRepoNameTagsGet) | **Get** /repositories/{repo_name}/tags | Get tags of a relevant repository.
[**RepositoriesRepoNameTagsPost**](ProductsApi.md#RepositoriesRepoNameTagsPost) | **Post** /repositories/{repo_name}/tags | Retag an image
[**RepositoriesRepoNameTagsTagDelete**](ProductsApi.md#RepositoriesRepoNameTagsTagDelete) | **Delete** /repositories/{repo_name}/tags/{tag} | Delete a tag in a repository.
[**RepositoriesRepoNameTagsTagGet**](ProductsApi.md#RepositoriesRepoNameTagsTagGet) | **Get** /repositories/{repo_name}/tags/{tag} | Get the tag of the repository.
[**RepositoriesRepoNameTagsTagLabelsGet**](ProductsApi.md#RepositoriesRepoNameTagsTagLabelsGet) | **Get** /repositories/{repo_name}/tags/{tag}/labels | Get labels of an image.
[**RepositoriesRepoNameTagsTagLabelsLabelIdDelete**](ProductsApi.md#RepositoriesRepoNameTagsTagLabelsLabelIdDelete) | **Delete** /repositories/{repo_name}/tags/{tag}/labels/{label_id} | Delete label from the image.
[**RepositoriesRepoNameTagsTagLabelsPost**](ProductsApi.md#RepositoriesRepoNameTagsTagLabelsPost) | **Post** /repositories/{repo_name}/tags/{tag}/labels | Add a label to image.
[**RepositoriesRepoNameTagsTagManifestGet**](ProductsApi.md#RepositoriesRepoNameTagsTagManifestGet) | **Get** /repositories/{repo_name}/tags/{tag}/manifest | Get manifests of a relevant repository.
[**RepositoriesRepoNameTagsTagScanPost**](ProductsApi.md#RepositoriesRepoNameTagsTagScanPost) | **Post** /repositories/{repo_name}/tags/{tag}/scan | Scan the image.
[**RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**](ProductsApi.md#RepositoriesRepoNameTagsTagVulnerabilityDetailsGet) | **Get** /repositories/{repo_name}/tags/{tag}/vulnerability/details | Get vulnerability details of the image.
[**RepositoriesTopGet**](ProductsApi.md#RepositoriesTopGet) | **Get** /repositories/top | Get public repositories which are accessed most.
[**SearchGet**](ProductsApi.md#SearchGet) | **Get** /search | Search for projects, repositories and helm charts
[**StatisticsGet**](ProductsApi.md#StatisticsGet) | **Get** /statistics | Get projects number and repositories number relevant to the user
[**SystemGcGet**](ProductsApi.md#SystemGcGet) | **Get** /system/gc | Get gc results.
[**SystemGcIdGet**](ProductsApi.md#SystemGcIdGet) | **Get** /system/gc/{id} | Get gc status.
[**SystemGcIdLogGet**](ProductsApi.md#SystemGcIdLogGet) | **Get** /system/gc/{id}/log | Get gc job log.
[**SystemGcScheduleGet**](ProductsApi.md#SystemGcScheduleGet) | **Get** /system/gc/schedule | Get gc&#39;s schedule.
[**SystemGcSchedulePost**](ProductsApi.md#SystemGcSchedulePost) | **Post** /system/gc/schedule | Create a gc schedule.
[**SystemGcSchedulePut**](ProductsApi.md#SystemGcSchedulePut) | **Put** /system/gc/schedule | Update gc&#39;s schedule.
[**SystemScanAllScheduleGet**](ProductsApi.md#SystemScanAllScheduleGet) | **Get** /system/scanAll/schedule | Get scan_all&#39;s schedule.
[**SystemScanAllSchedulePost**](ProductsApi.md#SystemScanAllSchedulePost) | **Post** /system/scanAll/schedule | Create a schedule or a manual trigger for the scan all job.
[**SystemScanAllSchedulePut**](ProductsApi.md#SystemScanAllSchedulePut) | **Put** /system/scanAll/schedule | Update scan all&#39;s schedule.
[**SysteminfoGet**](ProductsApi.md#SysteminfoGet) | **Get** /systeminfo | Get general system info
[**SysteminfoGetcertGet**](ProductsApi.md#SysteminfoGetcertGet) | **Get** /systeminfo/getcert | Get default root certificate.
[**SysteminfoVolumesGet**](ProductsApi.md#SysteminfoVolumesGet) | **Get** /systeminfo/volumes | Get system volume info (total/free size).
[**UsergroupsGet**](ProductsApi.md#UsergroupsGet) | **Get** /usergroups | Get all user groups information
[**UsergroupsGroupIdDelete**](ProductsApi.md#UsergroupsGroupIdDelete) | **Delete** /usergroups/{group_id} | Delete user group
[**UsergroupsGroupIdGet**](ProductsApi.md#UsergroupsGroupIdGet) | **Get** /usergroups/{group_id} | Get user group information
[**UsergroupsGroupIdPut**](ProductsApi.md#UsergroupsGroupIdPut) | **Put** /usergroups/{group_id} | Update group information
[**UsergroupsPost**](ProductsApi.md#UsergroupsPost) | **Post** /usergroups | Create user group
[**UsersCurrentGet**](ProductsApi.md#UsersCurrentGet) | **Get** /users/current | Get current user info.
[**UsersCurrentPermissionsGet**](ProductsApi.md#UsersCurrentPermissionsGet) | **Get** /users/current/permissions | Get current user permissions.
[**UsersGet**](ProductsApi.md#UsersGet) | **Get** /users | Get registered users of Harbor.
[**UsersPost**](ProductsApi.md#UsersPost) | **Post** /users | Creates a new user account.
[**UsersSearchGet**](ProductsApi.md#UsersSearchGet) | **Get** /users/search | Search users by username, email
[**UsersUserIdDelete**](ProductsApi.md#UsersUserIdDelete) | **Delete** /users/{user_id} | Mark a registered user as be removed.
[**UsersUserIdGenCliSecretPost**](ProductsApi.md#UsersUserIdGenCliSecretPost) | **Post** /users/{user_id}/gen_cli_secret | Generate new CLI secret for a user.
[**UsersUserIdGet**](ProductsApi.md#UsersUserIdGet) | **Get** /users/{user_id} | Get a user&#39;s profile.
[**UsersUserIdPasswordPut**](ProductsApi.md#UsersUserIdPasswordPut) | **Put** /users/{user_id}/password | Change the password on a user that already exists.
[**UsersUserIdPut**](ProductsApi.md#UsersUserIdPut) | **Put** /users/{user_id} | Update a registered user to change his profile.
[**UsersUserIdSysadminPut**](ProductsApi.md#UsersUserIdSysadminPut) | **Put** /users/{user_id}/sysadmin | Update a registered user to change to be an administrator of Harbor.


# **ChartrepoChartsPost**
> ChartrepoChartsPost(ctx, chart, optional)
Upload a chart file to the defult 'library' project.

Upload a chart file to the default 'library' project. Uploading together with the prov file at the same time is also supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **chart** | ***os.File**| The chart file | 
 **optional** | ***ChartrepoChartsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartrepoChartsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prov** | **optional.Interface of *os.File**| The provance file | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoHealthGet**
> InlineResponse2001 ChartrepoHealthGet(ctx, )
Check the health of chart repository service.

Check the health of chart repository service.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsGet**
> []ChartInfoEntry ChartrepoRepoChartsGet(ctx, repo)
Get all the charts under the specified project

Get all the charts under the specified project

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 

### Return type

[**[]ChartInfoEntry**](ChartInfoEntry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameDelete**
> ChartrepoRepoChartsNameDelete(ctx, repo, name)
Delete all the versions of the specified chart

Delete all the versions of the specified chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameGet**
> ChartrepoRepoChartsNameGet(ctx, repo, name)
Get all the versions of the specified chart

Get all the versions of the specified chart

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionDelete**
> ChartrepoRepoChartsNameVersionDelete(ctx, repo, name, version)
Delete the specified chart version

Delete the specified chart version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionGet**
> ChartrepoRepoChartsNameVersionGet(ctx, repo, name, version)
Get the specified chart version

Get the specified chart version

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionLabelsGet**
> ChartrepoRepoChartsNameVersionLabelsGet(ctx, repo, name, version)
Return the attahced labels of chart.

Return the attahced labels of the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionLabelsIdDelete**
> ChartrepoRepoChartsNameVersionLabelsIdDelete(ctx, repo, name, version, id)
Remove label from chart.

Remove label from the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 
  **id** | **int32**| The label ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsNameVersionLabelsPost**
> ChartrepoRepoChartsNameVersionLabelsPost(ctx, repo, name, version, label)
Mark label to chart.

Mark label to the specified chart version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **name** | **string**| The chart name | 
  **version** | **string**| The chart version | 
  **label** | [**Label**](Label.md)| The label being marked to the chart version | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoChartsPost**
> ChartrepoRepoChartsPost(ctx, repo, chart, optional)
Upload a chart file to the specified project.

Upload a chart file to the specified project. With this API, the corresponding provance file can be uploaded together with chart file at once.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **chart** | ***os.File**| The chart file | 
 **optional** | ***ChartrepoRepoChartsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChartrepoRepoChartsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **prov** | **optional.Interface of *os.File**| The provance file | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ChartrepoRepoProvPost**
> ChartrepoRepoProvPost(ctx, repo, prov)
Upload a provance file to the specified project.

Upload a provance file to the specified project. The provance file should be targeted for an existing chart file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repo** | **string**| The project name | 
  **prov** | ***os.File**| The provance file | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsGet**
> ConfigurationsResponse ConfigurationsGet(ctx, )
Get system configurations.

This endpoint is for retrieving system configurations that only provides for admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConfigurationsResponse**](ConfigurationsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ConfigurationsPut**
> ConfigurationsPut(ctx, configurations)
Modify system configurations.

This endpoint is for modifying system configurations that only provides for admin user. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configurations** | [**Configurations**](Configurations.md)| The configuration map can contain a subset of the attributes of the schema, which are to be updated. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EmailPingPost**
> EmailPingPost(ctx, optional)
Test connection and authentication with email server.

Test connection and authentication with email server. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmailPingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmailPingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **settings** | [**optional.Interface of EmailServerSetting**](EmailServerSetting.md)| Email server settings, if some of the settings are not assigned, they will be read from system configuration. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **HealthGet**
> OverallHealthStatus HealthGet(ctx, )
Health check API

The endpoint returns the health stauts of the system. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**OverallHealthStatus**](OverallHealthStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InternalSyncregistryPost**
> InternalSyncregistryPost(ctx, )
Sync repositories from registry to DB.

This endpoint is for syncing all repositories of registry with database. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsGet**
> []Label LabelsGet(ctx, scope, optional)
List labels according to the query strings.

This endpoint let user list labels by name, scope and project_id 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scope** | **string**| The label scope. Valid values are g and p. g for global labels and p for project labels. | 
 **optional** | ***LabelsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LabelsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The label name. | 
 **projectId** | **optional.Int64**| Relevant project ID, required when scope is p. | 
 **page** | **optional.Int32**| The page nubmer. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdDelete**
> LabelsIdDelete(ctx, id)
Delete the label specified by ID.

Delete the label specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdGet**
> Label LabelsIdGet(ctx, id)
Get the label specified by ID.

This endpoint let user get the label by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 

### Return type

[**Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdPut**
> LabelsIdPut(ctx, id, label)
Update the label properties.

This endpoint let user update label properties. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 
  **label** | [**Label**](Label.md)| The updated label json object. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsIdResourcesGet**
> Resource LabelsIdResourcesGet(ctx, id)
Get the resources that the label is referenced by.

This endpoint let user get the resources that the label is referenced by. Only the replication policies are returned for now. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label ID | 

### Return type

[**Resource**](Resource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LabelsPost**
> LabelsPost(ctx, label)
Post creates a label

This endpoint let user creates a label. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **label** | [**Label**](Label.md)| The json object of label. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapGroupsSearchGet**
> []UserGroup LdapGroupsSearchGet(ctx, optional)
Search available ldap groups.

This endpoint searches the available ldap groups based on related configuration parameters. support to search by groupname or groupdn. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapGroupsSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapGroupsSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupname** | **optional.String**| Ldap group name | 
 **groupdn** | **optional.String**| The LDAP group DN | 

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapPingPost**
> LdapPingPost(ctx, optional)
Ping available ldap service.

This endpoint ping the available ldap service for test related configuration parameters. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapPingPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapPingPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ldapconf** | [**optional.Interface of LdapConf**](LdapConf.md)| ldap configuration. support input ldap service configuration. If it&#39;s a empty request, will load current configuration from the system. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersImportPost**
> LdapUsersImportPost(ctx, uidList)
Import selected available ldap users.

This endpoint adds the selected available ldap users to harbor based on related configuration parameters from the system. System will try to guess the user email address and realname, add to harbor user information. If have errors when import user, will return the list of importing failed uid and the failed reason. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uidList** | [**LdapImportUsers**](LdapImportUsers.md)| The uid listed for importing. This list will check users validity of ldap service based on configuration from the system. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LdapUsersSearchGet**
> []LdapUsers LdapUsersSearchGet(ctx, optional)
Search available ldap users.

This endpoint searches the available ldap users based on related configuration parameters. Support searched by input ladp configuration, load configuration from the system and specific filter. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LdapUsersSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LdapUsersSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Registered user ID | 

### Return type

[**[]LdapUsers**](LdapUsers.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LogsGet**
> []AccessLog LogsGet(ctx, optional)
Get recent logs of the projects which the user is a member of

This endpoint let user see the recent operation logs of the projects which he is member of 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LogsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username of the operator. | 
 **repository** | **optional.String**| The name of repository | 
 **tag** | **optional.String**| The name of tag | 
 **operation** | **optional.String**| The operation | 
 **beginTimestamp** | **optional.String**| The begin timestamp | 
 **endTimestamp** | **optional.String**| The end timestamp | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsGet**
> []Project ProjectsGet(ctx, optional)
List projects

This endpoint returns all projects created by Harbor, and can be filtered by project name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The name of project. | 
 **public** | **optional.Bool**| The project is public or private. | 
 **owner** | **optional.String**| The name of project owner. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Project**](Project.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsHead**
> ProjectsHead(ctx, projectName)
Check if the project name user provided already exists.

This endpoint is used to check if the project name user provided already exist. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectName** | **string**| Project name for checking exists. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsPost**
> ProjectsPost(ctx, project)
Create a new project.

This endpoint is for user to create a new project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **project** | [**ProjectReq**](ProjectReq.md)| New created project. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdDelete**
> ProjectsProjectIdDelete(ctx, projectId)
Delete project by projectID

This endpoint is aimed to delete project by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID of project which will be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdGet**
> Project ProjectsProjectIdGet(ctx, projectId)
Return specific project detail information

This endpoint returns specific project information by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID for filtering results. | 

### Return type

[**Project**](Project.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdLogsGet**
> []AccessLog ProjectsProjectIdLogsGet(ctx, projectId, optional)
Get access logs accompany with a relevant project.

This endpoint let user search access logs filtered by operations and date time ranges. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID | 
 **optional** | ***ProjectsProjectIdLogsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsProjectIdLogsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **optional.String**| Username of the operator. | 
 **repository** | **optional.String**| The name of repository | 
 **tag** | **optional.String**| The name of tag | 
 **operation** | **optional.String**| The operation | 
 **beginTimestamp** | **optional.String**| The begin timestamp | 
 **endTimestamp** | **optional.String**| The end timestamp | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]AccessLog**](AccessLog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersGet**
> []ProjectMemberEntity ProjectsProjectIdMembersGet(ctx, projectId, optional)
Get all project member information

Get all project member information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
 **optional** | ***ProjectsProjectIdMembersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsProjectIdMembersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entityname** | **optional.String**| The entity name to search. | 

### Return type

[**[]ProjectMemberEntity**](ProjectMemberEntity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersMidDelete**
> ProjectsProjectIdMembersMidDelete(ctx, projectId, mid)
Delete project member

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **mid** | **int64**| Member ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersMidGet**
> ProjectMemberEntity ProjectsProjectIdMembersMidGet(ctx, projectId, mid)
Get the project member information

Get the project member information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **mid** | **int64**| The member ID | 

### Return type

[**ProjectMemberEntity**](ProjectMemberEntity.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersMidPut**
> ProjectsProjectIdMembersMidPut(ctx, projectId, mid, optional)
Update project member

Update project member relationship

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
  **mid** | **int64**| Member ID. | 
 **optional** | ***ProjectsProjectIdMembersMidPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsProjectIdMembersMidPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | [**optional.Interface of RoleRequest**](RoleRequest.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMembersPost**
> ProjectsProjectIdMembersPost(ctx, projectId, optional)
Create project member

Create project member relationship, the member can be one of the user_member and group_member,  The user_member need to specify user_id or username. If the user already exist in harbor DB, specify the user_id,  If does not exist in harbor DB, it will SearchAndOnBoard the user. The group_member need to specify id or ldap_group_dn. If the group already exist in harbor DB. specify the user group's id,  If does not exist, it will SearchAndOnBoard the group. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Relevant project ID. | 
 **optional** | ***ProjectsProjectIdMembersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectsProjectIdMembersPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectMember** | [**optional.Interface of ProjectMember**](ProjectMember.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasGet**
> ProjectMetadata ProjectsProjectIdMetadatasGet(ctx, projectId)
Get project metadata.

This endpoint returns metadata of the project specified by project ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 

### Return type

[**ProjectMetadata**](ProjectMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNameDelete**
> ProjectsProjectIdMetadatasMetaNameDelete(ctx, projectId, metaName)
Delete metadata of a project

This endpoint is aimed to delete metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNameGet**
> ProjectMetadata ProjectsProjectIdMetadatasMetaNameGet(ctx, projectId, metaName)
Get project metadata

This endpoint returns specified metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Project ID for filtering results. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

[**ProjectMetadata**](ProjectMetadata.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasMetaNamePut**
> ProjectsProjectIdMetadatasMetaNamePut(ctx, projectId, metaName)
Update metadata of a project.

This endpoint is aimed to update the metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| The ID of project. | 
  **metaName** | **string**| The name of metadat. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdMetadatasPost**
> ProjectsProjectIdMetadatasPost(ctx, projectId, metadata)
Add metadata for the project.

This endpoint is aimed to add metadata of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Selected project ID. | 
  **metadata** | [**ProjectMetadata**](ProjectMetadata.md)| The metadata of project. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ProjectsProjectIdPut**
> ProjectsProjectIdPut(ctx, projectId, project)
Update properties for a selected project.

This endpoint is aimed to update the properties of a project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int64**| Selected project ID. | 
  **project** | [**ProjectReq**](ProjectReq.md)| Updates of project. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

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

# **RegistriesGet**
> []Registry RegistriesGet(ctx, optional)
List registries.

This endpoint let user list filtered registries by name, if name is nil, list returns all registries. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RegistriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Registry&#39;s name. | 

### Return type

[**[]Registry**](Registry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdDelete**
> RegistriesIdDelete(ctx, id)
Delete specific registry.

This endpoint is for to delete specific registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry&#39;s ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdGet**
> Registry RegistriesIdGet(ctx, id)
Get registry.

This endpoint is for get specific registry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry ID. | 

### Return type

[**Registry**](Registry.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdInfoGet**
> RegistryInfo RegistriesIdInfoGet(ctx, id)
Get registry info.

Get the info of one specific registry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry ID. | 

### Return type

[**RegistryInfo**](RegistryInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdNamespaceGet**
> []Namespace RegistriesIdNamespaceGet(ctx, id, optional)
List namespaces of registry

This endpoint let user list namespaces of registry according to query. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| The registry ID. | 
 **optional** | ***RegistriesIdNamespaceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RegistriesIdNamespaceGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| The name of namespace. | 

### Return type

[**[]Namespace**](Namespace.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesIdPut**
> RegistriesIdPut(ctx, id, repoTarget)
Update a given registry.

This endpoint is for update a given registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The registry&#39;s ID. | 
  **repoTarget** | [**PutRegistry**](PutRegistry.md)| Updates registry. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesPingPost**
> RegistriesPingPost(ctx, registry)
Ping status of a registry.

This endpoint checks status of a registry, the registry can be given by ID or URL (together with credential) 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registry** | [**Registry**](Registry.md)| Registry to ping. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistriesPost**
> RegistriesPost(ctx, registry)
Create a new registry.

This endpoint is for user to create a new registry. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **registry** | [**Registry**](Registry.md)| New created registry. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationAdaptersGet**
> []string ReplicationAdaptersGet(ctx, )
List supported adapters.

This endpoint let user list supported adapters. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsGet**
> []ReplicationExecution ReplicationExecutionsGet(ctx, optional)
List replication executions.

This endpoint let user list replication executions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationExecutionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationExecutionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **optional.Int32**| The policy ID. | 
 **status** | **optional.String**| The execution status. | 
 **trigger** | **optional.String**| The trigger mode. | 
 **page** | **optional.Int32**| The page. | 
 **pageSize** | **optional.Int32**| The page size. | 

### Return type

[**[]ReplicationExecution**](ReplicationExecution.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdGet**
> ReplicationExecution ReplicationExecutionsIdGet(ctx, id)
Get the execution of the replication.

This endpoint is for user to get one execution of the replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 

### Return type

[**ReplicationExecution**](ReplicationExecution.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdPut**
> ReplicationExecutionsIdPut(ctx, id)
Stop the execution of the replication.

This endpoint is for user to stop one execution of the replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdTasksGet**
> []ReplicationTask ReplicationExecutionsIdTasksGet(ctx, id)
Get the task list of one execution.

This endpoint is for user to get the task list of one execution. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 

### Return type

[**[]ReplicationTask**](ReplicationTask.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsIdTasksTaskIdLogGet**
> ReplicationExecutionsIdTasksTaskIdLogGet(ctx, id, taskId)
Get the log of one task.

This endpoint is for user to get the log of one task. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The execution ID. | 
  **taskId** | **int64**| The task ID. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationExecutionsPost**
> ReplicationExecutionsPost(ctx, execution)
Start one execution of the replication.

This endpoint is for user to start one execution of the replication. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **execution** | [**ReplicationExecution**](ReplicationExecution.md)| The execution that needs to be started, only the property \&quot;policy_id\&quot; is needed. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesGet**
> []ReplicationPolicy ReplicationPoliciesGet(ctx, optional)
List replication policies

This endpoint let user list replication policies 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ReplicationPoliciesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReplicationPoliciesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| The replication policy name. | 
 **page** | **optional.Int32**| The page nubmer. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesIdDelete**
> ReplicationPoliciesIdDelete(ctx, id)
Delete the replication policy specified by ID.

Delete the replication policy specified by ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Replication policy ID | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesIdGet**
> ReplicationPolicy ReplicationPoliciesIdGet(ctx, id)
Get replication policy.

This endpoint let user get replication policy by specific ID. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| policy ID | 

### Return type

[**ReplicationPolicy**](ReplicationPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesIdPut**
> ReplicationPoliciesIdPut(ctx, id, policy)
Update the replication policy

This endpoint let user update policy. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| policy ID | 
  **policy** | [**ReplicationPolicy**](ReplicationPolicy.md)| The replication policy model. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReplicationPoliciesPost**
> ReplicationPoliciesPost(ctx, policy)
Create a replication policy

This endpoint let user create a replication policy 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **policy** | [**ReplicationPolicy**](ReplicationPolicy.md)| The policy model. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesGet**
> []Repository RepositoriesGet(ctx, projectId, optional)
Get repositories accompany with relevant project and repo name.

This endpoint lets user search repositories accompanying with relevant project ID and repo name. Repositories can be sorted by repo name, creation_time, update_time in either ascending or descending order. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectId** | **int32**| Relevant project ID. | 
 **optional** | ***RepositoriesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **optional.String**| Repo name for filtering results. | 
 **sort** | **optional.String**| Sort method, valid values include: &#39;name&#39;, &#39;-name&#39;, &#39;creation_time&#39;, &#39;-creation_time&#39;, &#39;update_time&#39;, &#39;-update_time&#39;. Here &#39;-&#39; stands for descending order.  | 
 **labelId** | **optional.Int32**| The ID of label used to filter the result. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page, default is 10, maximum is 100. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameDelete**
> RepositoriesRepoNameDelete(ctx, repoName)
Delete a repository.

This endpoint let user delete a repository with name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameLabelsGet**
> []Label RepositoriesRepoNameLabelsGet(ctx, repoName)
Get labels of a repository.

Get labels of a repository specified by the repo_name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameLabelsLabelIdDelete**
> RepositoriesRepoNameLabelsLabelIdDelete(ctx, repoName, labelId)
Delete label from the repository.

Delete the label from the repository specified by the repo_name. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **labelId** | **int32**| The ID of label. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameLabelsPost**
> RepositoriesRepoNameLabelsPost(ctx, repoName, label)
Add a label to the repository.

Add a label to the repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **label** | [**Label**](Label.md)| Only the ID property is required. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNamePut**
> RepositoriesRepoNamePut(ctx, repoName, description)
Update description of the repository.

This endpoint is used to update description of the repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 
  **description** | [**RepositoryDescription**](RepositoryDescription.md)| The description of the repository. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameSignaturesGet**
> []RepoSignature RepositoriesRepoNameSignaturesGet(ctx, repoName)
Get signature information of a repository

This endpoint aims to retrieve signature information of a repository, the data is from the nested notary instance of Harbor. If the repository does not have any signature information in notary, this API will return an empty list with response code 200, instead of 404 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| repository name. | 

### Return type

[**[]RepoSignature**](RepoSignature.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsGet**
> []DetailedTag RepositoriesRepoNameTagsGet(ctx, repoName, optional)
Get tags of a relevant repository.

This endpoint aims to retrieve tags from a relevant repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 
 **optional** | ***RepositoriesRepoNameTagsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesRepoNameTagsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **labelIds** | **optional.String**| A list of comma separated label IDs. | 

### Return type

[**[]DetailedTag**](DetailedTag.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsPost**
> RepositoriesRepoNameTagsPost(ctx, repoName, request)
Retag an image

This endpoint tags an existing image with another tag in this repo, source images can be in different repos or projects. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 
  **request** | [**RetagReq**](RetagReq.md)| Request to give source image and target tag. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagDelete**
> RepositoriesRepoNameTagsTagDelete(ctx, repoName, tag)
Delete a tag in a repository.

This endpoint let user delete tags with repo name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository which will be deleted. | 
  **tag** | **string**| Tag of a repository. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagGet**
> DetailedTag RepositoriesRepoNameTagsTagGet(ctx, repoName, tag)
Get the tag of the repository.

This endpoint aims to retrieve the tag of the repository. If deployed with Notary, the signature property of response represents whether the image is singed or not. If the property is null, the image is unsigned. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Relevant repository name. | 
  **tag** | **string**| Tag of the repository. | 

### Return type

[**DetailedTag**](DetailedTag.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagLabelsGet**
> []Label RepositoriesRepoNameTagsTagLabelsGet(ctx, repoName, tag)
Get labels of an image.

Get labels of an image specified by the repo_name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **tag** | **string**| The tag of the image. | 

### Return type

[**[]Label**](Label.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagLabelsLabelIdDelete**
> RepositoriesRepoNameTagsTagLabelsLabelIdDelete(ctx, repoName, tag, labelId)
Delete label from the image.

Delete the label from the image specified by the repo_name and tag. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **tag** | **string**| The tag of the image. | 
  **labelId** | **int32**| The ID of label. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagLabelsPost**
> RepositoriesRepoNameTagsTagLabelsPost(ctx, repoName, tag, label)
Add a label to image.

Add a label to the image. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| The name of repository. | 
  **tag** | **string**| The tag of the image. | 
  **label** | [**Label**](Label.md)| Only the ID property is required. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagManifestGet**
> Manifest RepositoriesRepoNameTagsTagManifestGet(ctx, repoName, tag, optional)
Get manifests of a relevant repository.

This endpoint aims to retreive manifests from a relevant repository. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 
 **optional** | ***RepositoriesRepoNameTagsTagManifestGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesRepoNameTagsTagManifestGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **version** | **optional.String**| The version of manifest, valid value are \&quot;v1\&quot; and \&quot;v2\&quot;, default is \&quot;v2\&quot; | 

### Return type

[**Manifest**](Manifest.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagScanPost**
> RepositoriesRepoNameTagsTagScanPost(ctx, repoName, tag)
Scan the image.

Trigger jobservice to call Clair API to scan the image identified by the repo_name and tag.  Only project admins have permission to scan images under the project. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesRepoNameTagsTagVulnerabilityDetailsGet**
> []VulnerabilityItem RepositoriesRepoNameTagsTagVulnerabilityDetailsGet(ctx, repoName, tag)
Get vulnerability details of the image.

Call Clair API to get the vulnerability based on the previous successful scan. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **repoName** | **string**| Repository name | 
  **tag** | **string**| Tag name | 

### Return type

[**[]VulnerabilityItem**](VulnerabilityItem.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RepositoriesTopGet**
> []Repository RepositoriesTopGet(ctx, optional)
Get public repositories which are accessed most.

This endpoint aims to let users see the most popular public repositories 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RepositoriesTopGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RepositoriesTopGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **count** | **optional.Int32**| The number of the requested public repositories, default is 10 if not provided. | 

### Return type

[**[]Repository**](Repository.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchGet**
> []Search SearchGet(ctx, q)
Search for projects, repositories and helm charts

The Search endpoint returns information about the projects ,repositories  and helm charts offered at public status or related to the current logged in user. The response includes the project, repository list and charts in a proper display order. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **q** | **string**| Search parameter for project and repository name. | 

### Return type

[**[]Search**](Search.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StatisticsGet**
> StatisticMap StatisticsGet(ctx, )
Get projects number and repositories number relevant to the user

This endpoint is aimed to statistic all of the projects number and repositories number relevant to the logined user, also the public projects number and repositories number. If the user is admin, he can also get total projects number and total repositories number. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**StatisticMap**](StatisticMap.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcGet**
> []GcResult SystemGcGet(ctx, )
Get gc results.

This endpoint let user get latest ten gc results.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]GcResult**](GCResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcIdGet**
> GcResult SystemGcIdGet(ctx, id)
Get gc status.

This endpoint let user get gc status filtered by specific ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Relevant job ID | 

### Return type

[**GcResult**](GCResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcIdLogGet**
> string SystemGcIdLogGet(ctx, id)
Get gc job log.

This endpoint let user get gc job logs filtered by specific ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Relevant job ID | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcScheduleGet**
> AdminJobSchedule SystemGcScheduleGet(ctx, )
Get gc's schedule.

This endpoint is for get schedule of gc job.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdminJobSchedule**](AdminJobSchedule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcSchedulePost**
> SystemGcSchedulePost(ctx, schedule)
Create a gc schedule.

This endpoint is for update gc schedule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**AdminJobSchedule**](AdminJobSchedule.md)| Updates of gc&#39;s schedule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemGcSchedulePut**
> SystemGcSchedulePut(ctx, schedule)
Update gc's schedule.

This endpoint is for update gc schedule. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**AdminJobSchedule**](AdminJobSchedule.md)| Updates of gc&#39;s schedule. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemScanAllScheduleGet**
> AdminJobSchedule SystemScanAllScheduleGet(ctx, )
Get scan_all's schedule.

This endpoint is for getting a schedule for the scan all job, which scans all of images in Harbor.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AdminJobSchedule**](AdminJobSchedule.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemScanAllSchedulePost**
> SystemScanAllSchedulePost(ctx, schedule)
Create a schedule or a manual trigger for the scan all job.

This endpoint is for creating a schedule or a manual trigger for the scan all job, which scans all of images in Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**AdminJobSchedule**](AdminJobSchedule.md)| Create a schedule or a manual trigger for the scan all job. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemScanAllSchedulePut**
> SystemScanAllSchedulePut(ctx, schedule)
Update scan all's schedule.

This endpoint is for updating the schedule of scan all job, which scans all of images in Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **schedule** | [**AdminJobSchedule**](AdminJobSchedule.md)| Updates the schedule of scan all job, which scans all of images in Harbor. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGet**
> GeneralInfo SysteminfoGet(ctx, )
Get general system info

This API is for retrieving general system info, this can be called by anonymous request. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GeneralInfo**](GeneralInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoGetcertGet**
> SysteminfoGetcertGet(ctx, )
Get default root certificate.

This endpoint is for downloading a default root certificate. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SysteminfoVolumesGet**
> SystemInfo SysteminfoVolumesGet(ctx, )
Get system volume info (total/free size).

This endpoint is for retrieving system volume info that only provides for admin user. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SystemInfo**](SystemInfo.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGet**
> []UserGroup UsergroupsGet(ctx, )
Get all user groups information

Get all user groups information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserGroup**](UserGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGroupIdDelete**
> UsergroupsGroupIdDelete(ctx, groupId)
Delete user group

Delete user group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int32**|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGroupIdGet**
> UserGroup UsergroupsGroupIdGet(ctx, groupId)
Get user group information

Get user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 

### Return type

[**UserGroup**](UserGroup.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsGroupIdPut**
> UsergroupsGroupIdPut(ctx, groupId, optional)
Update group information

Update user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **groupId** | **int64**| Group ID | 
 **optional** | ***UsergroupsGroupIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsGroupIdPutOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **usergroup** | [**optional.Interface of UserGroup**](UserGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsergroupsPost**
> UsergroupsPost(ctx, optional)
Create user group

Create user group information

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsergroupsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsergroupsPostOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usergroup** | [**optional.Interface of UserGroup**](UserGroup.md)|  | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersCurrentGet**
> User UsersCurrentGet(ctx, )
Get current user info.

This endpoint is to get the current user information. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersCurrentPermissionsGet**
> []Permission UsersCurrentPermissionsGet(ctx, optional)
Get current user permissions.

This endpoint is to get the current user permissions. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersCurrentPermissionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersCurrentPermissionsGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **optional.String**| Get permissions of the scope | 
 **relative** | **optional.Bool**| If true, the resources in the response are relative to the scope, eg for resource &#39;/project/1/repository&#39; if relative is &#39;true&#39; then the resource in response will be &#39;repository&#39;.  | 

### Return type

[**[]Permission**](Permission.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersGet**
> []User UsersGet(ctx, optional)
Get registered users of Harbor.

This endpoint is for user to search registered users, support for filtering results with username.Notice, by now this operation is only for administrator. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username for filtering results. | 
 **email** | **optional.String**| Email for filtering results. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersPost**
> UsersPost(ctx, user)
Creates a new user account.

This endpoint is to create a user if the user does not already exist. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **user** | [**User**](User.md)| New created user. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSearchGet**
> []UserSearch UsersSearchGet(ctx, optional)
Search users by username, email

This endpoint is to search the users by username, email. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UsersSearchGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UsersSearchGetOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **username** | **optional.String**| Username for filtering results. | 
 **email** | **optional.String**| Email for filtering results. | 
 **page** | **optional.Int32**| The page nubmer, default is 1. | 
 **pageSize** | **optional.Int32**| The size of per page. | 

### Return type

[**[]UserSearch**](UserSearch.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdDelete**
> UsersUserIdDelete(ctx, userId)
Mark a registered user as be removed.

This endpoint let administrator of Harbor mark a registered user as be removed.It actually won't be deleted from DB. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| User ID for marking as to be removed. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdGenCliSecretPost**
> InlineResponse200 UsersUserIdGenCliSecretPost(ctx, userId)
Generate new CLI secret for a user.

This endpoint let user generate a new CLI secret for himself.  This API only works when auth mode is set to 'OIDC'. Once this API returns with successful status, the old secret will be invalid, as there will be only one CLI secret for a user.  The new secret will be returned in the response. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| User ID | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdGet**
> User UsersUserIdGet(ctx, userId)
Get a user's profile.

Get user's profile with user id. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 

### Return type

[**User**](User.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPasswordPut**
> UsersUserIdPasswordPut(ctx, userId, password)
Change the password on a user that already exists.

This endpoint is for user to update password. Users with the admin role can change any user's password. Guest users can change only their own password. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID. | 
  **password** | [**Password**](Password.md)| Password to be updated, the attribute &#39;old_password&#39; is optional when the API is called by the system administrator. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdPut**
> UsersUserIdPut(ctx, userId, profile)
Update a registered user to change his profile.

This endpoint let a registered user change his profile. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 
  **profile** | [**UserProfile**](UserProfile.md)| Only email, realname and comment can be modified. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersUserIdSysadminPut**
> UsersUserIdSysadminPut(ctx, userId, hasAdminRole)
Update a registered user to change to be an administrator of Harbor.

This endpoint let a registered user change to be an administrator of Harbor. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **int32**| Registered user ID | 
  **hasAdminRole** | [**HasAdminRole**](HasAdminRole.md)| Toggle a user to admin or not. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

