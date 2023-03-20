// Package response
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package response

import "time"

type YopResponse struct {
	metadata     *YopResponseMetadata
	result       any
	stringResult string
}

type YosUploadResponse struct {
	metadata     *YosUploadResponseMetadata
	result       any
	stringResult string
}

type YosDownloadResponse struct {
}

type YopResponseMetadata struct {
	yopRequestId    string
	yopSign         string
	contentType     string
	date            time.Time
	server          string
	yopCertSerialNo string
}

type YosUploadResponseMetadata struct {
	YopResponseMetadata
	crc64ECMA string
}
