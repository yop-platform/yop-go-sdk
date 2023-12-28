// Package response
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/22 10:54 PM
package response

import (
	"encoding/json"
	"errors"
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var ANALYZER_CHAIN = []HttpResponseAnalyzer{
	&YopMetadataResponseAnalyzer{},
	&YopSignatureCheckAnalyzer{},
	&YopErrorResponseAnalyzer{},
	&YopJsonResponseAnalyzer{},
}

type HttpResponseAnalyzer interface {
	Analyze(context RespHandleContext, httpResponse *http.Response) error
}

type YopMetadataResponseAnalyzer struct {
}

func (yopMetadataResponseAnalyzer *YopMetadataResponseAnalyzer) Analyze(context RespHandleContext, httpResponse *http.Response) error {
	var metadata = YopResponseMetadata{}
	metadata.YopRequestId = httpResponse.Header.Get(constants.YOP_REQUEST_ID)
	metadata.YopContentSha256 = httpResponse.Header.Get(constants.YOP_CONTENT_SHA256)
	metadata.YopSign = httpResponse.Header.Get(constants.YOP_SIGN)
	metadata.ContentType = httpResponse.Header.Get(constants.CONTENT_TYPE)
	d, _ := time.Parse(time.RFC1123, httpResponse.Header.Get(constants.DATE))
	metadata.Date = d
	metadata.YopCertSerialNo = httpResponse.Header.Get(constants.YOP_SIGN_CERT_SERIAL_NO)
	metadata.Crc64ECMA = httpResponse.Header.Get(constants.YOP_HASH_CRC64ECMA)
	context.YopResponse.Metadata = &metadata
	return nil
}

type YopSignatureCheckAnalyzer struct {
}

func (yopSignatureCheckAnalyzer *YopSignatureCheckAnalyzer) Analyze(context RespHandleContext, httpResponse *http.Response) error {
	signature := context.YopResponse.Result.Sign

	if 0 < len(signature) {
		b, err := json.Marshal(context.YopResponse.Result.Result)
		if nil != err {
			return errors.New("unexpected error")
		}
		if !context.YopSigner.VerifyResponse(string(b), signature, context.YopRequest.PlatformPubKey) {
			return errors.New("response sign verify failure")
		}
	}
	return nil
}

type YopErrorResponseAnalyzer struct {
}

func (yopErrorResponseAnalyzer *YopErrorResponseAnalyzer) Analyze(context RespHandleContext, httpResponse *http.Response) error {
	var statusCode = httpResponse.StatusCode
	log.Println("statusCode:" + strconv.Itoa(statusCode))
	if statusCode/100 == constants.SC_OK && statusCode != constants.SC_NO_CONTENT {
		return nil
	}
	var yopServiceError = YopServiceError{}
	json.Unmarshal(context.YopResponse.Content, &yopServiceError)
	if 0 < len(yopServiceError.Message) {
		return &yopServiceError
	}
	return nil
}

type YopJsonResponseAnalyzer struct {
}

func (yopJsonResponseAnalyzer *YopJsonResponseAnalyzer) Analyze(context RespHandleContext, httpResponse *http.Response) error {
	if 0 < len(context.YopResponse.Content) && strings.HasPrefix(context.YopResponse.Metadata.ContentType, constants.YOP_HTTP_CONTENT_TYPE_JSON) {
		json.Unmarshal(context.YopResponse.Content, &context.YopResponse)
	}
	return nil
}
