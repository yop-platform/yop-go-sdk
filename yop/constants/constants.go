// Package constants
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/21 3:24 PM
package constants

const (
	DEFAULT_YOP_PROTOCOL_VERSION  = "yop-auth-v3"
	DEFAULT_EXPIRATION_IN_SECONDS = 1800
	YOP_CONTENT_SHA256            = "x-yop-content-sha256"
	AUTHORIZATION                 = "Authorization"
	YOP_REQUEST_ID                = "x-yop-request-id"
	YOP_APPKEY_HEADER_KEY         = "x-yop-appkey"
	YOP_SIGN_HEADER_KEY           = "x-yop-sign"
	USER_AGENT_HEADER_KEY         = "User-Agent"
	CONTENT_TYPE                  = "Content-Type"
	DEFAULT_USER_AGENT            = ""
	YOP_SIGN                      = "x-yop-sign"
	YOP_SIGN_CERT_SERIAL_NO       = "x-yop-sign-serial-no"
	DATE                          = "Date"
	YOP_HASH_CRC64ECMA            = "x-yop-hash-crc64ecma"

	YOP_HTTP_CONTENT_TYPE_JSON           = "application/json"
	YOP_HTTP_CONTENT_TYPE_FORM           = "application/x-www-form-urlencoded;charset=utf-8"
	YOP_HTTP_CONTENT_TYPE_MULTIPART_FORM = "multipart/form-data"
	YOP_HTTP_CONTENT_TYPE_STREAM         = "application/octet-stream"
	YOP_HTTP_CONTENT_TYPE_TEXT           = "text/plain;charset=UTF-8"

	POST_HTTP_METHOD = "POST"
	GET_HTTP_METHOD  = "GET"

	SC_OK         = 200
	SC_NO_CONTENT = 204

	SDK_VERSION = "4.3.2"
)
