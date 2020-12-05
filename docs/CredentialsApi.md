# \CredentialsApi

All URIs are relative to *https://api.lucfish.com/storage/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Credentials**](CredentialsApi.md#Credentials) | **Post** /credentials | 获取上传凭证



## Credentials

> CredentialsRsp Credentials(ctx, optional)

获取上传凭证

获取上传凭证

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timestamp** | **optional.String**|  | 
 **noncestr** | **optional.String**|  | 
 **signature** | **optional.String**|  | 
 **credentialsReq** | [**optional.Interface of CredentialsReq**](CredentialsReq.md)|  | 

### Return type

[**CredentialsRsp**](CredentialsRsp.md)

### Authorization

[ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

