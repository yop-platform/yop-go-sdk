/*
 * @Author: BigRocs
 * @Date: 2024-10-27 08:51:30
 * @LastEditTime: 2024-10-27 08:52:23
 * @LastEditors: BigRocs
 * @Description: QQ: 532388887, Email:bigrocs@qq.com
 */
// Package auth
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/21 10:36 AM
package auth

import (
	"testing"

	"github.com/yop-platform/yop-go-sdk/yop/request"
)

func TestRsaSigner_SignRequest(t *testing.T) {
	var yopRequest = buildYopRequest()
	var signer = RsaSigner{}
	signer.SignRequest(*yopRequest)
	t.Log(yopRequest.Headers)
}

func buildYopRequest() *request.YopRequest {
	var isvPriKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDBHBdHbQXsPT+EpAhLA9k2Q5O8GLCAUFLWYB57Uhc4ZNa2YUhjrTFvFZMFQuMjaVgdmFGTvqfGYUQBRldHFhf9kuXf5LPb+m0BJ/R5AWCyTcX7DHouoGODMfxkCZrimILwYWDkwhYTHr5hEV58nGRQtHOIVB5a4i/y4Z1vvX1MIjA+8OJ3zpaxXKkj+46OtfmjloUPGFSzz+rqrRRtMqYePLkWZ0J+CmIXM1Kwl/kgYUq/YGSYy9Q5vTojN9WBKzk7euOoCcsWtRrBQysdyM3yDPXjhnXx7G0nh07hSUh+rMDZ7Zst0lVQol/7kPXzNwh6eUmRGY9lfruIMS5kg1ydAgMBAAECggEAD4yQf0rTCEOiQq7mkAu+SLVGRwYB6EMPeH2C1tE0V3EfLM5GgugmK9ij3u+U1HweATwLjYbzgXDBhgzA6FNqGRvj8JQ8u0C92DL8Z2XqAFFs2JsXl3uIp761oOR5GTfIi0x7/c928ZEvKSe54PTCyxDMoLSNQSonTDpIb//k/+U4xEOQ1mjlSvlOM5ic7/kdw+G+aP/Hk/T6kg/vIblWQHx8SB3WYpLb/R6oPO+05X+zcQ+vVX1TrQ/amDp6/PouWjTF5hf48JEBdM8+xJzUwnalrG9U7pChfyGAOXQT1fbDdywBJXt6pZsT/mz1RkUC5Uto5/aVQGIDD+IPm/ZDbQKBgQDEwF09sjUb5hHEdmG28RmMf4E0JCOEzCvxiUpovobymqapLM5bf2oLNXqGenEAMbfFQatJFVKx6YBZwFIj/xzQJt8fL/jRzlbLijaANP+1JacvTsfXKBXS888FN3rkKisTlhmYXI+4EwA1wbcRkLDH3vezdVCi9cszQ9HvwkVfOwKBgQD7QvzD3pirXIJ64JizWTS4MJMko3CWepsq9UZ5uyoHWh7tSz86H/2y0FK10YpJEJeGtyPXlnU+uQwjYMJRPLlNv8180pjCJX2ZTW2drB2vOJvormhMhDIYAZtPAHu2dajzdy4VRuvFTtH4FpW/KjAJrTLK3ze3K95ACYVBJ8EmBwKBgB0K8DiNN724hmLjvqTMjiLpJ19U/lE5+jqbM3qmtTDWl0ddr9BdzH9/E2kKZefLbv8VJH2TQjO07hdRhk599/jZ5BGseSQvOyysaEMgj6ZjunwHOwSNjDspdiOk/uTzPIyVmY2eDDD1zRAiWi2jmBTI2vOIm7CSa75TgofLu4XFAoGBAJrFM4+vYNlFXbY0/LqU+21ttmV+K471rPj0Jto7GPN4Zs6CaEr0g8COpDQNA6JoDv5Td0eIDWZ6c+ii5G9H+VjUCc6WprQIhepVkGzsJUjWlOrp66MeVwEElFdAk/PbXBvEUOWYTwi1uY6Y0trzMK31OvFOODKjWf6WHrf4tfgnAoGAOri6bX2D/zqpJT3mJ5MIVJJbn4D4Idx+TCUaVRSY1rBp+Y2ofW1W8ktu7xPO9/LwVQR7kJeosEBAFGTmGqll033ywu5+8X8J1bw6HCghkI0yHW752sOdfl30kXi3Ds8tQsvSEHRfnPb8yvWve2srZb9ubwOvpI0PtOIujZP4fYI="
	var priKey = &request.IsvPriKey{Value: isvPriKey, CertType: "RSA2048"}
	var result = request.BuildYopRequest()
	result.AppId = "app_1234567"
	result.ApiUri = "/rest/v2.0/yop/platform/certs"
	result.ServerRoot = "http://ycetest.yeepay.com:30228/yop-center"
	result.HttpMethod = "GET"
	result.IsvPriKey = *priKey
	result.AddParam("name", "testName")
	result.AddParam("age", 18)
	return result
}
