# \CredentialsApi

All URIs are relative to *https://api.xres.lucfish.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](CredentialsApi.md#Create) | **Post** /oss/credentials | 获取上传凭证 credentials



## Create

> CreateUploadCredentialsRsp Create(ctx, timestamp, noncestr, signature, body)

获取上传凭证 credentials

获取上传凭证 credentials

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timestamp** | **string**| timestamp | 
**noncestr** | **string**| noncestr | 
**signature** | **string**| signature | 
**body** | [**CreateUploadCredentialsReq**](CreateUploadCredentialsReq.md)|  | 

### Return type

[**CreateUploadCredentialsRsp**](CreateUploadCredentialsRsp.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

