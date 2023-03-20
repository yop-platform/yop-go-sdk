// Package client
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package client

import (
	"github.com/yop-platform/yop-go-sdk/yop/request"
	response "github.com/yop-platform/yop-go-sdk/yop/response"
)

type YopClient interface {
	request(yopRequest request.YopRequest) response.YopResponse

	download(yopRequest request.YopRequest) response.YosDownloadResponse

	upload(yopRequest request.YopRequest) response.YosUploadResponse
}
