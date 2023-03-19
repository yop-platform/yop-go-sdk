// Package request
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package request

import (
	"go/types"
	"strings"
)

const (
	SERVER_ROOT     = "http://openapi.yeepay.com/yop-center"
	YOS_SERVER_ROOT = "http://yos.yeepay.com/yop-center"
)

type YopRequest struct {
	// 服务地址，一般情况无需指定
	serverRoot string
	apiUri     string
	httpMethod string
	appId      string
	priKey     *IsvPriKey
	// form请求的参数
	params types.Map
	// json请求参数
	content string
}

type IsvPriKey struct {
	// 密钥类型：RSA2048
	certType string
	// 私钥值
	value string
}

func (request *YopRequest) setServerRoot(serverRoot string) {
	request.serverRoot = serverRoot
}

func (request *YopRequest) handleServerRoot() {
	if 0 != len(request.serverRoot) {
		return
	}

	if 0 == len(request.apiUri) || !strings.HasPrefix(request.apiUri, "/yos") {
		request.serverRoot = SERVER_ROOT
	}
	request.serverRoot = YOS_SERVER_ROOT

}
