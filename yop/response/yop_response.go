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
	Result   any
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
