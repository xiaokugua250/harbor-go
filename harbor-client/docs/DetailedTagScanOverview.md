# DetailedTagScanOverview

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** | The digest of the image. | [optional] [default to null]
**ScanStatus** | **string** | The status of the scan job, it can be \&quot;pendnig\&quot;, \&quot;running\&quot;, \&quot;finished\&quot;, \&quot;error\&quot;. | [optional] [default to null]
**JobId** | **int32** | The ID of the job on jobservice to scan the image. | [optional] [default to null]
**Severity** | **int32** | 0-Not scanned, 1-Negligible, 2-Unknown, 3-Low, 4-Medium, 5-High | [optional] [default to null]
**DetailsKey** | **string** | The top layer name of this image in Clair, this is for calling Clair API to get the vulnerability list of this image. | [optional] [default to null]
**Components** | [***DetailedTagScanOverviewComponents**](DetailedTag_scan_overview_components.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


