/*
 * storage open api
 *
 * storage open api
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package storage
// CredentialsReq 获取上传凭证
type CredentialsReq struct {
	// 上传渠道，腾讯云：qcloud、阿里云：aliyun、七牛：qiniu
	Channel string `json:"channel"`
	// 计算凭证参数，具体数据结构与上传渠道对应
	Params map[string]interface{} `json:"params"`
}
