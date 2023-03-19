// Package client
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package client

import (
	response "github.com/yop-platform/yop-go-sdk/yop/reqponse"
	"github.com/yop-platform/yop-go-sdk/yop/request"
)

type YopClient interface {
	request(yopRequest request.YopRequest) response.YopResponse

	download(yopRequest request.YopRequest) response.YosDownloadResponse

	upload(yopRequest request.YopRequest) response.YosUploadResponse
}
