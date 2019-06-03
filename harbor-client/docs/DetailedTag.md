# DetailedTag

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** | The digest of the tag. | [optional] [default to null]
**Name** | **string** | The name of the tag. | [optional] [default to null]
**Size** | **int32** | The size of the image. | [optional] [default to null]
**Architecture** | **string** | The architecture of the image. | [optional] [default to null]
**Os** | **string** | The os of the image. | [optional] [default to null]
**DockerVersion** | **string** | The version of docker which builds the image. | [optional] [default to null]
**Author** | **string** | The author of the image. | [optional] [default to null]
**Created** | **string** | The build time of the image. | [optional] [default to null]
**Signature** | [***interface{}**](interface{}.md) | The signature of image, defined by RepoSignature. If it is null, the image is unsigned. | [optional] [default to null]
**ScanOverview** | [***DetailedTagScanOverview**](DetailedTag_scan_overview.md) |  | [optional] [default to null]
**Labels** | [**[]Label**](Label.md) | The label list. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


