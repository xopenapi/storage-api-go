# Go API client for storage

storage open api

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:

```golang
import "./storage"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.xres.lucfish.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CredentialsApi* | [**Create**](docs/CredentialsApi.md#create) | **Post** /oss/credentials | 获取上传凭证 credentials


## Documentation For Models

 - [CreateUploadCredentialsReq](docs/CreateUploadCredentialsReq.md)
 - [CreateUploadCredentialsRsp](docs/CreateUploadCredentialsRsp.md)


## Documentation For Authorization

 Endpoints do not require authorization.



## Author



