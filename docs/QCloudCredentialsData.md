# QCloudCredentialsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionToken** | **string** | 请求时需要用的 token 字符串，最终请求 COS API 时，需要放在 Header 的 x-cos-security-token 字段 | 
**TmpSecretId** | **string** | 临时密钥 Id，可用于计算签名 | 
**TmpSecretKey** | **string** | 临时密钥 Key，可用于计算签名 | 
**StartTime** | **int64** | 密钥的起始时间戳（毫秒） | 
**ExpiredTime** | **int64** | 密钥的失效时间戳（毫秒） | 
**RequestId** | **string** | 请求ID | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


