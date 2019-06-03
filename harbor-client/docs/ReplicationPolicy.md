# ReplicationPolicy

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The policy ID. | [optional] [default to null]
**Name** | **string** | The policy name. | [optional] [default to null]
**Description** | **string** | The description of the policy. | [optional] [default to null]
**SrcRegistry** | [***Registry**](Registry.md) | The source registry. | [optional] [default to null]
**DestRegistry** | [***Registry**](Registry.md) | The destination registry. | [optional] [default to null]
**DestNamespace** | **string** | The destination namespace. | [optional] [default to null]
**Trigger** | [***ReplicationTrigger**](ReplicationTrigger.md) |  | [optional] [default to null]
**Filters** | [**[]ReplicationFilter**](ReplicationFilter.md) | The replication policy filter array. | [optional] [default to null]
**Deletion** | **bool** | Whether to replicate the deletion operation. | [optional] [default to null]
**Override** | **bool** | Whether to override the resources on the destination registry. | [optional] [default to null]
**Enabled** | **bool** | Whether the policy is enabled or not. | [optional] [default to null]
**CreationTime** | **string** | The create time of the policy. | [optional] [default to null]
**UpdateTime** | **string** | The update time of the policy. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


