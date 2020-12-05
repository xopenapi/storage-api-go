/*
 * storage open api
 *
 * storage open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package storage-api-go
// CredentialsRsp 获取上传凭证结果
type CredentialsRsp struct {
	// 错误码
	Code int32 `json:"code"`
	// 错误消息
	Msg string `json:"msg,omitempty"`
	// 上传凭证数据
	Data OneOfQCloudCredentialsData `json:"data,omitempty"`
}
