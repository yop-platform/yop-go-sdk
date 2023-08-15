// Package client
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package client

import (
	"bytes"
	"context"
	"errors"
	"github.com/gofrs/uuid/v5"
	"github.com/yop-platform/yop-go-sdk/yop/auth"
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"github.com/yop-platform/yop-go-sdk/yop/response"
	"github.com/yop-platform/yop-go-sdk/yop/utils"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

var DefaultClient = YopClient{&http.Client{Transport: http.DefaultTransport}}

type YopClient struct {
	*http.Client
}

// Request 普通请求
func (yopClient *YopClient) Request(request *request.YopRequest) (*response.YopResponse, error) {
	initRequest(request)
	var signer = auth.RsaSigner{}
	err := signer.SignRequest(*request)
	if nil != err {
		return nil, err
	}

	httpRequest, err := buildHttpRequest(*request)
	if nil != err {
		return nil, err
	}
	httpResp, err := yopClient.Client.Do(&httpRequest)
	if nil != err {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if nil != err {
		return nil, err
	}
	var yopResponse = response.YopResponse{Content: body}
	metaData := response.YopResponseMetadata{}
	metaData.YopSign = httpResp.Header.Get("X-Yop-Sign")
	metaData.YopRequestId = httpResp.Header.Get("X-Yop-Request-Id")
	yopResponse.Metadata = &metaData
	context := response.RespHandleContext{YopSigner: &signer, YopResponse: &yopResponse, YopRequest: *request}
	for i := range response.ANALYZER_CHAIN {
		err = response.ANALYZER_CHAIN[i].Analyze(context, httpResp)
		if nil != err {
			return nil, err
		}
	}
	return &yopResponse, nil
}
func initRequest(yopRequest *request.YopRequest) {
	yopRequest.RequestId = uuid.Must(uuid.NewV4()).String()
	log.Println("requestId:" + yopRequest.RequestId)
	if 0 == len(yopRequest.ServerRoot) {
		yopRequest.HandleServerRoot()
	}
	if 0 == len(yopRequest.PlatformPubKey.Value) {
		yopRequest.PlatformPubKey.Value = request.YOP_PLATFORM_PUBLIC_KEY
		yopRequest.PlatformPubKey.CertType = request.RSA2048
	}
	addStandardHeaders(yopRequest)
}
func addStandardHeaders(yopRequest *request.YopRequest) {
	yopRequest.Headers = map[string]string{}
	yopRequest.Headers[constants.YOP_REQUEST_ID] = yopRequest.RequestId
	yopRequest.Headers[constants.YOP_APPKEY_HEADER_KEY] = yopRequest.AppId
	yopRequest.Headers[constants.USER_AGENT_HEADER_KEY] = buildUserAgent()
}

func buildUserAgent() string {
	return "go" + "/" + constants.SDK_VERSION + "/" + runtime.GOOS + "/" + runtime.Version()
}

func buildHttpRequest(yopRequest request.YopRequest) (http.Request, error) {
	if yopRequest.Timeout == 0 {
		yopRequest.Timeout = 10 * time.Second
	}
	ctx, _ := context.WithTimeout(context.Background(), yopRequest.Timeout)
	//defer cancel()

	var uri = yopRequest.ServerRoot + yopRequest.ApiUri
	isMultiPart, err := checkForMultiPart(yopRequest)
	if nil != err {
		return http.Request{}, err
	}
	var result http.Request
	if isMultiPart {
		bodyBuf := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuf)

		for k, v := range yopRequest.Params {
			for i := range v {
				bodyWriter.WriteField(k, url.QueryEscape(v[i]))
			}
		}

		for k, v := range yopRequest.Files {
			fileWriter, _ := bodyWriter.CreateFormFile(k, v.Name())
			io.Copy(fileWriter, v)
		}
		bodyWriter.Close()

		if err != nil {
			return http.Request{}, err
		}
		req, err := http.NewRequestWithContext(ctx, "POST", uri, bodyBuf)
		if nil != err {
			return http.Request{}, err
		}
		req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
		result = *req
	} else {
		var encodedParam = utils.EncodeParameters(yopRequest.Params, false)
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
						formValues.Set(k, url.QueryEscape(v[i]))
					}
				}
				formDataStr := formValues.Encode()
				body = bytes.NewBuffer([]byte(formDataStr))
			}
		}
		httpRequest, err := http.NewRequestWithContext(ctx, yopRequest.HttpMethod, uri, body)
		if err != nil {
			return http.Request{}, err
		}
		result = *httpRequest
		result.Header.Set(constants.CONTENT_TYPE, getContentType(yopRequest))
	}
	for k, v := range yopRequest.Headers {
		result.Header.Set(k, v)
	}
	return result, err
}

func checkForMultiPart(yopRequest request.YopRequest) (bool, error) {
	var result = nil != yopRequest.Files && 0 < len(yopRequest.Files)
	if result && 0 != strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod) {
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
