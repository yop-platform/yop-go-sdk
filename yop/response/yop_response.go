// Package response
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package response

import (
	"github.com/yop-platform/yop-go-sdk/yop/auth"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"time"
)

type YopResponse struct {
	Metadata *YopResponseMetadata
	Result   YopResult
	// http请求收到的原始响应体
	// 若接口类型为文件下载，则该值为文件内容
	Content []byte
}

type YopResponseMetadata struct {
	YopRequestId     string
	YopContentSha256 string
	YopSign          string
	ContentType      string
	Date             time.Time
	Server           string
	YopCertSerialNo  string
	Crc64ECMA        string
}

type RespHandleContext struct {
	auth.YopSigner
	*YopResponse
	request.YopRequest
}

type YopResult struct {
	// 业务处理状态
	State string `json:"state"`
	// 业务响应结果
	Result any `json:"result"`
	// 错误信息
	Error YopError `json:"error"`
	// 时间戳
	TS int64 `json:"ts"`
	// 签名
	Sign string `json:"sign"`
}

type YopError struct {
	// 错误码
	Code string `json:"code"`
	// 错误信息
	Message string `json:"message"`
	// 请求标识
	Solution string `json:"solution"`
}
