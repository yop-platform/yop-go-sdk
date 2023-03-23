// Package client
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/21 7:47 PM
package client

import (
	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"github.com/yop-platform/yop-go-sdk/yop/response"
	"os"
	"testing"
)

func TestYopClient_GET_Request(t *testing.T) {
	yopResp, err := DefaultClient.Request(*buildGetYopRequest())
	testAssert(yopResp, err, t)
}

func TestYopClient_Post_Json_Request(t *testing.T) {
	yopResp, err := DefaultClient.Request(*buildJsonYopRequest())
	testAssert(yopResp, err, t)
}

func TestYopClient_Post_From_Request(t *testing.T) {
	yopResp, err := DefaultClient.Request(*buildPostFormYopRequest())
	testAssert(yopResp, err, t)
}

func TestYopClient_Upload(t *testing.T) {
	yopResp, err := DefaultClient.Request(*buildUploadYopRequest())
	testAssert(yopResp, err, t)
}

func testAssert(resp *response.YopResponse, err error, t *testing.T) {
	if nil != err {
		t.Fatal(err.Error())
	}
	t.Log(resp.Result)
}

var (
	isvPriKey      = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDBHBdHbQXsPT+EpAhLA9k2Q5O8GLCAUFLWYB57Uhc4ZNa2YUhjrTFvFZMFQuMjaVgdmFGTvqfGYUQBRldHFhf9kuXf5LPb+m0BJ/R5AWCyTcX7DHouoGODMfxkCZrimILwYWDkwhYTHr5hEV58nGRQtHOIVB5a4i/y4Z1vvX1MIjA+8OJ3zpaxXKkj+46OtfmjloUPGFSzz+rqrRRtMqYePLkWZ0J+CmIXM1Kwl/kgYUq/YGSYy9Q5vTojN9WBKzk7euOoCcsWtRrBQysdyM3yDPXjhnXx7G0nh07hSUh+rMDZ7Zst0lVQol/7kPXzNwh6eUmRGY9lfruIMS5kg1ydAgMBAAECggEAD4yQf0rTCEOiQq7mkAu+SLVGRwYB6EMPeH2C1tE0V3EfLM5GgugmK9ij3u+U1HweATwLjYbzgXDBhgzA6FNqGRvj8JQ8u0C92DL8Z2XqAFFs2JsXl3uIp761oOR5GTfIi0x7/c928ZEvKSe54PTCyxDMoLSNQSonTDpIb//k/+U4xEOQ1mjlSvlOM5ic7/kdw+G+aP/Hk/T6kg/vIblWQHx8SB3WYpLb/R6oPO+05X+zcQ+vVX1TrQ/amDp6/PouWjTF5hf48JEBdM8+xJzUwnalrG9U7pChfyGAOXQT1fbDdywBJXt6pZsT/mz1RkUC5Uto5/aVQGIDD+IPm/ZDbQKBgQDEwF09sjUb5hHEdmG28RmMf4E0JCOEzCvxiUpovobymqapLM5bf2oLNXqGenEAMbfFQatJFVKx6YBZwFIj/xzQJt8fL/jRzlbLijaANP+1JacvTsfXKBXS888FN3rkKisTlhmYXI+4EwA1wbcRkLDH3vezdVCi9cszQ9HvwkVfOwKBgQD7QvzD3pirXIJ64JizWTS4MJMko3CWepsq9UZ5uyoHWh7tSz86H/2y0FK10YpJEJeGtyPXlnU+uQwjYMJRPLlNv8180pjCJX2ZTW2drB2vOJvormhMhDIYAZtPAHu2dajzdy4VRuvFTtH4FpW/KjAJrTLK3ze3K95ACYVBJ8EmBwKBgB0K8DiNN724hmLjvqTMjiLpJ19U/lE5+jqbM3qmtTDWl0ddr9BdzH9/E2kKZefLbv8VJH2TQjO07hdRhk599/jZ5BGseSQvOyysaEMgj6ZjunwHOwSNjDspdiOk/uTzPIyVmY2eDDD1zRAiWi2jmBTI2vOIm7CSa75TgofLu4XFAoGBAJrFM4+vYNlFXbY0/LqU+21ttmV+K471rPj0Jto7GPN4Zs6CaEr0g8COpDQNA6JoDv5Td0eIDWZ6c+ii5G9H+VjUCc6WprQIhepVkGzsJUjWlOrp66MeVwEElFdAk/PbXBvEUOWYTwi1uY6Y0trzMK31OvFOODKjWf6WHrf4tfgnAoGAOri6bX2D/zqpJT3mJ5MIVJJbn4D4Idx+TCUaVRSY1rBp+Y2ofW1W8ktu7xPO9/LwVQR7kJeosEBAFGTmGqll033ywu5+8X8J1bw6HCghkI0yHW752sOdfl30kXi3Ds8tQsvSEHRfnPb8yvWve2srZb9ubwOvpI0PtOIujZP4fYI="
	platformPubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4g7dPL+CBeuzFmARI2GFjZpKODUROaMG+E6wdNfv5lhPqC3jjTIeljWU8AiruZLGRhl92QWcTjb3XonjaV6k9rf9adQtyv2FLS7bl2Vz2WgjJ0FJ5/qMaoXaT+oAgWFk2GypyvoIZsscsGpUStm6BxpWZpbPrGJR0N95un/130cQI9VCmfvgkkCaXt7TU1BbiYzkc8MDpLScGm/GUCB2wB5PclvOxvf5BR/zNVYywTEFmw2Jo0hIPPSWB5Yyf2mx950Fx8da56co/FxLdMwkDOO51Qg3fbaExQDVzTm8Odi++wVJEP1y34tlmpwFUVbAKIEbyyELmi/2S6GG0j9vNwIDAQAB"
)

func buildGetYopRequest() *request.YopRequest {
	var priKey = &request.IsvPriKey{Value: isvPriKey, CertType: request.RSA2048}
	var yopRequest = request.BuildYopRequest()
	var platformPub = &request.PlatformPubKey{Value: platformPubKey, CertType: request.RSA2048}
	yopRequest.PlatformPubKey = platformPub
	yopRequest.AppId = "app_15958159879157110001"
	yopRequest.ApiUri = "/rest/v1.0/test-wdc/product-query/query-for-doc"
	yopRequest.ServerRoot = "http://ycetest.yeepay.com:30228/yop-center"
	yopRequest.HttpMethod = constants.GET_HTTP_METHOD
	yopRequest.IsvPriKey = priKey
	yopRequest.AddParam("string0", "le1")
	return yopRequest
}

func buildJsonYopRequest() *request.YopRequest {
	var priKey = &request.IsvPriKey{Value: isvPriKey, CertType: request.RSA2048}
	var result = request.BuildYopRequest()
	var platformPub = &request.PlatformPubKey{Value: platformPubKey, CertType: request.RSA2048}
	result.PlatformPubKey = platformPub
	result.AppId = "app_15958159879157110001"
	result.ApiUri = "/rest/v1.0/kj/transferdomestic/singlequery"
	result.ServerRoot = "http://ycetest.yeepay.com:30228/yop-center"
	result.HttpMethod = constants.POST_HTTP_METHOD
	result.IsvPriKey = priKey
	result.Content = "{\"merchantId\":\"1595815987915711\",\"requestId\":\"requestId\"}"
	return result
}

func buildPostFormYopRequest() *request.YopRequest {
	var priKey = &request.IsvPriKey{Value: isvPriKey, CertType: request.RSA2048}
	var result = request.BuildYopRequest()
	var platformPub = &request.PlatformPubKey{Value: platformPubKey, CertType: request.RSA2048}
	result.PlatformPubKey = platformPub
	result.AppId = "app_15958159879157110001"
	result.ApiUri = "/rest/v1.0/wym/test/notify"
	result.ServerRoot = "http://ycetest.yeepay.com:30228/yop-center"
	result.HttpMethod = constants.POST_HTTP_METHOD
	result.IsvPriKey = priKey
	result.AddParam("string", "testString")
	return result
}

func buildUploadYopRequest() *request.YopRequest {
	var priKey = &request.IsvPriKey{Value: isvPriKey, CertType: request.RSA2048}
	var result = request.BuildYopRequest()
	var platformPub = &request.PlatformPubKey{Value: platformPubKey, CertType: request.RSA2048}
	result.PlatformPubKey = platformPub
	result.AppId = "app_15958159879157110001"
	result.ApiUri = "/yos/v1.0/test/upload"
	result.ServerRoot = "http://ycetest.yeepay.com:30228/yop-center"
	result.HttpMethod = constants.POST_HTTP_METHOD
	result.IsvPriKey = priKey
	var path = "/Users/yp-21024/go/src/yop-go-sdk/README.md"
	f, _ := os.Open(path)
	result.AddFile("file", f)
	result.AddParam("string", "ppp")
	return result
}
