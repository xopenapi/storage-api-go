/*
 * Storage open api
 *
 * storage open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package storage
// CreateUploadCredentialsRsp struct for CreateUploadCredentialsRsp
type CreateUploadCredentialsRsp struct {
	Code int32 `json:"code,omitempty"`
	Msg string `json:"msg,omitempty"`
	Data map[string]interface{} `json:"data,omitempty"`
}
