// Package client
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package client

import (
	"bytes"
	"errors"
	"github.com/yop-platform/yop-go-sdk/yop/auth"
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"github.com/yop-platform/yop-go-sdk/yop/response"
	"github.com/yop-platform/yop-go-sdk/yop/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

var DefaultClient = YopClient{&http.Client{Transport: http.DefaultTransport}}

type YopClient struct {
	*http.Client
}

var signer = auth.RsaSigner{}

// Request 普通请求
func (yopClient *YopClient) Request(request *request.YopRequest) (*response.YopResponse, error) {
	log.Println("requestId:" + request.RequestId)
	addStandardHeaders(request)

	signer.SignRequest(request)

	httpRequest, err := buildHttpRequest(*request)
	if nil != err {
		return nil, err
	}
	resp, _ := yopClient.Client.Do(&httpRequest)

	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
	return nil, nil
}

// Upload 文件上传
func (*YopClient) Upload(request *request.YopRequest) *response.YopResponse {
	return nil
}

// Download 文件下载
func (*YopClient) Download(request *request.YopRequest) *response.YopResponse {
	return nil
}

func addStandardHeaders(yopRequest *request.YopRequest) {
	yopRequest.Headers[constants.YOP_REQUEST_ID] = yopRequest.RequestId
	yopRequest.Headers[constants.YOP_APPKEY_HEADER_KEY] = yopRequest.AppId
	yopRequest.Headers[constants.USER_AGENT_HEADER_KEY] = buildUserAgent()
}

func buildUserAgent() string {
	return "go" + "/" + "4.0.0" + "/" + runtime.GOOS + "/" + runtime.Version() + runtime.GOROOT()
}

func buildHttpRequest(yopRequest request.YopRequest) (http.Request, error) {
	var err error = nil
	var uri = yopRequest.ServerRoot + yopRequest.ApiUri
	isMultiPart, err := checkForMultiPart(yopRequest)
	if nil != err {
		return http.Request{}, err
	}
	var result http.Request
	if isMultiPart {

	} else {
		var encodedParam = utils.EncodeParameters(yopRequest.Params)
		var requestHasPayload = 0 < len(yopRequest.Content)
		var requestIsPost = 0 == strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod)
		var putParamsInUri = !requestIsPost || requestHasPayload
		if 0 < len(encodedParam) && putParamsInUri {
			uri += "?" + encodedParam
		}
		var body io.Reader = nil
		if 0 == strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod) {
			if 0 < len(yopRequest.Content) {
				body = bytes.NewBuffer([]byte(yopRequest.Content))
			} else {
				formValues := url.Values{}
				for k, v := range yopRequest.Params {
					for i := range v {
						formValues.Set(k, v[i])
					}
				}
				formDataStr := formValues.Encode()
				body = bytes.NewBuffer([]byte(formDataStr))
			}
		}
		httpRequest, _ := http.NewRequest(yopRequest.HttpMethod, uri, body)
		result = *httpRequest
	}
	for k, v := range yopRequest.Headers {
		result.Header.Set(k, v)
	}
	result.Header.Set(constants.CONTENT_TYPE, getContentType(yopRequest))
	return result, err
}

func checkForMultiPart(yopRequest request.YopRequest) (bool, error) {
	var result = nil != yopRequest.Files && 0 < len(yopRequest.Files)
	if result && 0 == strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod) {
		var errorMsg = "ContentType:multipart/form-data only support Post Request"
		log.Fatal(errorMsg)
		return false, errors.New(errorMsg)
	}
	return result, nil
}

func getContentType(yopRequest request.YopRequest) string {
	if 0 == strings.Compare("POST", yopRequest.HttpMethod) && 0 < len(yopRequest.Content) {
		return constants.YOP_HTTP_CONTENT_TYPE_JSON
	}
	if 0 < len(yopRequest.Params) {
		return constants.YOP_HTTP_CONTENT_TYPE_FORM
	}
	return constants.YOP_HTTP_CONTENT_TYPE_FORM
}
