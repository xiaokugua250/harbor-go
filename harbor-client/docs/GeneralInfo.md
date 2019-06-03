# GeneralInfo

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WithNotary** | **bool** | If the Harbor instance is deployed with nested notary. | [optional] [default to null]
**WithClair** | **bool** | If the Harbor instance is deployed with nested clair. | [optional] [default to null]
**WithAdmiral** | **bool** | If the Harbor instance is deployed with Admiral. | [optional] [default to null]
**AdmiralEndpoint** | **string** | The url of the endpoint of admiral instance. | [optional] [default to null]
**RegistryUrl** | **string** | The url of registry against which the docker command should be issued. | [optional] [default to null]
**ExternalUrl** | **string** | The external URL of Harbor, with protocol. | [optional] [default to null]
**AuthMode** | **string** | The auth mode of current Harbor instance. | [optional] [default to null]
**ProjectCreationRestriction** | **string** | Indicate who can create projects, it could be &#39;adminonly&#39; or &#39;everyone&#39;. | [optional] [default to null]
**SelfRegistration** | **bool** | Indicate whether the Harbor instance enable user to register himself. | [optional] [default to null]
**HasCaRoot** | **bool** | Indicate whether there is a ca root cert file ready for download in the file system. | [optional] [default to null]
**HarborVersion** | **string** | The build version of Harbor. | [optional] [default to null]
**NextScanAll** | **int32** | The UTC time in milliseconds, after which user can call scanAll API to scan all images. | [optional] [default to null]
**ClairVulnerabilityStatus** | [***GeneralInfoClairVulnerabilityStatus**](GeneralInfo_clair_vulnerability_status.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


