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
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/yop-platform/yop-go-sdk/yop/auth"
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"github.com/yop-platform/yop-go-sdk/yop/response"
	"github.com/yop-platform/yop-go-sdk/yop/utils"
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

	// 设置超时时间，如果没有设置则使用默认值
	if request.Timeout == 0 {
		request.Timeout = 10 * time.Second
	}

	// 创建带超时的 context，在整个请求过程中保持有效
	ctx, cancel := context.WithTimeout(context.Background(), request.Timeout)
	defer cancel()

	httpRequest, err := buildHttpRequestWithContext(ctx, *request)
	if nil != err {
		return nil, err
	}
	httpResp, err := yopClient.Client.Do(&httpRequest)
	if nil != err {
		return nil, err
	}
	defer func() {
		if closeErr := httpResp.Body.Close(); closeErr != nil {
			utils.Logger.Warnf("Failed to close response body: %v", closeErr)
		}
	}()
	body, err := io.ReadAll(httpResp.Body)
	if nil != err {
		return nil, err
	}
	var yopResponse = response.YopResponse{Content: body}
	metaData := response.YopResponseMetadata{}
	metaData.YopSign = httpResp.Header.Get("X-Yop-Sign")
	metaData.YopRequestId = httpResp.Header.Get("X-Yop-Request-Id")
	yopResponse.Metadata = &metaData
	handleContext := response.RespHandleContext{YopSigner: &signer, YopResponse: &yopResponse, YopRequest: *request}
	for i := range response.ANALYZER_CHAIN {
		err = response.ANALYZER_CHAIN[i].Analyze(handleContext, httpResp)
		if nil != err {
			return nil, err
		}
	}
	return &yopResponse, nil
}

func initRequest(yopRequest *request.YopRequest) {
	yopRequest.RequestId = utils.GenerateRequestID()
	utils.Logger.Info("requestId:" + yopRequest.RequestId)
	if len(yopRequest.ServerRoot) == 0 {
		yopRequest.HandleServerRoot()
	}
	if len(yopRequest.PlatformPubKey.Value) == 0 {
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

// buildHttpRequestWithContext 使用外部传入的 context 构建 HTTP 请求
func buildHttpRequestWithContext(ctx context.Context, yopRequest request.YopRequest) (http.Request, error) {
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
				if err := bodyWriter.WriteField(k, url.QueryEscape(v[i])); err != nil {
					return http.Request{}, err
				}
			}
		}

		for k, v := range yopRequest.Files {
			// 只使用文件名，避免路径遍历安全问题
			fileName := filepath.Base(v.Name())
			fileWriter, err := bodyWriter.CreateFormFile(k, fileName)
			if err != nil {
				return http.Request{}, err
			}
			if _, err := io.Copy(fileWriter, v); err != nil {
				return http.Request{}, err
			}
		}
		if err := bodyWriter.Close(); err != nil {
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
		var requestHasPayload = len(yopRequest.Content) > 0
		var requestIsPost = strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod) == 0
		var putParamsInUri = !requestIsPost || requestHasPayload
		if len(encodedParam) > 0 && putParamsInUri {
			uri += "?" + encodedParam
		}
		var body io.Reader = nil
		if strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod) == 0 {
			if len(yopRequest.Content) > 0 {
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
	var result = len(yopRequest.Files) > 0
	if result && strings.Compare(constants.POST_HTTP_METHOD, yopRequest.HttpMethod) != 0 {
		var errorMsg = "ContentType:multipart/form-data only support Post Request"
		utils.Logger.Error("error: " + errorMsg)
		return false, errors.New(errorMsg)
	}
	return result, nil
}

func getContentType(yopRequest request.YopRequest) string {
	if strings.Compare("POST", yopRequest.HttpMethod) == 0 && len(yopRequest.Content) > 0 {
		return constants.YOP_HTTP_CONTENT_TYPE_JSON
	}
	if len(yopRequest.Params) > 0 {
		return constants.YOP_HTTP_CONTENT_TYPE_FORM
	}
	return constants.YOP_HTTP_CONTENT_TYPE_FORM
}
