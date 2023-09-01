# yop-go-sdk
Yeepay openapi SDK fo Go
## 本项目的使用场景
1.请求易宝开放平台接口 </br>
2.解密易宝开放平台回调内容</br>
3.构造易宝收银台签名
## 使用
1.引入
```go
"github.com/yop-platform/yop-go-sdk/yop/client"
"github.com/yop-platform/yop-go-sdk/yop/constants"
"github.com/yop-platform/yop-go-sdk/yop/request"
"github.com/yop-platform/yop-go-sdk/yop/utils"
```
2.调用接口</br>
放在最前面：请求超时时间通过YopRequest.Timeout设置，默认值为10秒<br><br>
Get请求
```go

var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.GET_HTTP_METHOD,"/rest/v1.0/test-wdc/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{
	// request failed
}
//yopResp.Result为请求结果
```
Post Form请求
```go

var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test-wdc/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
	// request failed
}
//yopResp.Result为请求结果
```
Post Json请求
```go
  var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
  var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD,"/rest/v1.0/test-wdc/product-query/query-for-doc")
  yopRequest.AppId = "appId"
  yopRequest.IsvPriKey = priKey
  // 设置json请求报文
  var params = map[string]any{}
  params["merchantId"] = "1595815987915711"
  params["requestId"] = "requestId"
  result.Content = utils.ParseToJsonStr(params)
  
  yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
	// request failed
}
  //yopResp.Result为请求结果
```
文件上传请求
```go

var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD,"/rest/v1.0/test-wdc/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
result.AddFile("file", f)
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
	// request failed
}
// yopResp.Result为上传请求结果
```
文件下载请求
```go
  var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
  var yopRequest = request.NewYopRequest(constants.GET_HTTP_METHOD,"/rest/v1.0/test-wdc/product-query/query-for-doc")
  yopRequest.AppId = "appId"
  yopRequest.IsvPriKey = priKey
  yopRequest.AddParam("paramName", "paramValue")
  yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
	// request failed
}
  // yopResp.Content为文件内容
```
解密回调内容
```go
  //utils.DecryptCallback

  var callback = "Ars6jASSiylO70_VJDQ5SFU1zQwaI36kG5WhlSKHjkGdU3fEVEkkbhvAxKjOTUiw9vF7RMnmGKQQWAuV8jCKaOpMNjIEMHehBaPASwTiEE946CcbOeoNILGHf0o20xj2gqqvkQToFXEMNiic7bcYbfi0PxIrR6loBZnW-m5bqzB5RXLibiSjGlmr5CDnxV4tZXmYlkkeN2BcT4msWjfCtuaTMK_fN77WJcCMlW7ffqiN5yIOeqB4QBb5lOnClTRW4DThKPOMkXupAM2AnPxTkDp4n9lh-SK56zLuafk1bQhWUNcS9L4YEKZGJIjP7DY20TAWEr3yXo8w0w0VtB13Ig$Xf6fETKWcLTudBh2HluGSQTqhBRJa6EXHhXlMryWW8Y384RjVwIfpQm19RmTgkoqRc2tNcTWxRIW6itIS62DrzixlqRa099jx21uGqt8FCpvdWwnwlC16SgkeU_5NnrpjA_WQ0XW9RhNxzuQmwfxHGbtnth4vNXWswcSm23j3KQaXFjVP5Ws1uYVCxYSLMxqJE7a56DNWONGcGJJsc0KTCc7cdfr8n24emAaPCNteIG2RM8F17pRxY5yVnguTSZPXmhBlyI25xS7rciWzKZLp2Kfh_JCivABbA-_5Vf3VWPmjITs-TR5HlGVFbnT0eOUMUepXUemjjP8R0f8cBeH2NKej6QjQL99tvlrrxg_QfmezE0WTCITCNDBhpbHiq90lFyLjwlWNDTRo8rhjouSlMA9Ae_b-B4eZorDRVxw3BWywdyo2FzNk-dUDeBVaIth9YsaMGsq9XivGjlnnx3YEVfEtuVSvEm1xBdYsTHcM02nMwZb8Ze2WL1kIFo8IFM0$AES$SHA256"
  // 此处为测试数据，正式使用是请使用真实的平台公钥
  var platformPubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7LqdMV7ZeOWUwVp0duSucTr4VwUNHtYLlWEUWlBtDQDEPhx0WZZdw2DxEbQqMQM5BjXZACYlhEdPt0HicDthOIUeUt8JNcvgq06vIE958RzgVBa5z3zvMLYWJIZaUyxsxC7Us06eNiB+du0rEBxUckru41ZSu/DX9jssFC+l5459b3WWELNf2fXqJyfb4f8GuGk8enXgJdxBUcmwgaEQxJjWkPqhzSiRy9GKjcXBdCkzCYR4xmLkHe6K0YFiBxax7lOni3zVOsvHC9XdhbepwB9fMkHbZXS/LJf5aS5ltendObpVrAD9kck7bIQzsrM49/SG/dYmbtm139I6ygsCzQIDAQAB"
  var isvPriKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDBHBdHbQXsPT+EpAhLA9k2Q5O8GLCAUFLWYB57Uhc4ZNa2YUhjrTFvFZMFQuMjaVgdmFGTvqfGYUQBRldHFhf9kuXf5LPb+m0BJ/R5AWCyTcX7DHouoGODMfxkCZrimILwYWDkwhYTHr5hEV58nGRQtHOIVB5a4i/y4Z1vvX1MIjA+8OJ3zpaxXKkj+46OtfmjloUPGFSzz+rqrRRtMqYePLkWZ0J+CmIXM1Kwl/kgYUq/YGSYy9Q5vTojN9WBKzk7euOoCcsWtRrBQysdyM3yDPXjhnXx7G0nh07hSUh+rMDZ7Zst0lVQol/7kPXzNwh6eUmRGY9lfruIMS5kg1ydAgMBAAECggEAD4yQf0rTCEOiQq7mkAu+SLVGRwYB6EMPeH2C1tE0V3EfLM5GgugmK9ij3u+U1HweATwLjYbzgXDBhgzA6FNqGRvj8JQ8u0C92DL8Z2XqAFFs2JsXl3uIp761oOR5GTfIi0x7/c928ZEvKSe54PTCyxDMoLSNQSonTDpIb//k/+U4xEOQ1mjlSvlOM5ic7/kdw+G+aP/Hk/T6kg/vIblWQHx8SB3WYpLb/R6oPO+05X+zcQ+vVX1TrQ/amDp6/PouWjTF5hf48JEBdM8+xJzUwnalrG9U7pChfyGAOXQT1fbDdywBJXt6pZsT/mz1RkUC5Uto5/aVQGIDD+IPm/ZDbQKBgQDEwF09sjUb5hHEdmG28RmMf4E0JCOEzCvxiUpovobymqapLM5bf2oLNXqGenEAMbfFQatJFVKx6YBZwFIj/xzQJt8fL/jRzlbLijaANP+1JacvTsfXKBXS888FN3rkKisTlhmYXI+4EwA1wbcRkLDH3vezdVCi9cszQ9HvwkVfOwKBgQD7QvzD3pirXIJ64JizWTS4MJMko3CWepsq9UZ5uyoHWh7tSz86H/2y0FK10YpJEJeGtyPXlnU+uQwjYMJRPLlNv8180pjCJX2ZTW2drB2vOJvormhMhDIYAZtPAHu2dajzdy4VRuvFTtH4FpW/KjAJrTLK3ze3K95ACYVBJ8EmBwKBgB0K8DiNN724hmLjvqTMjiLpJ19U/lE5+jqbM3qmtTDWl0ddr9BdzH9/E2kKZefLbv8VJH2TQjO07hdRhk599/jZ5BGseSQvOyysaEMgj6ZjunwHOwSNjDspdiOk/uTzPIyVmY2eDDD1zRAiWi2jmBTI2vOIm7CSa75TgofLu4XFAoGBAJrFM4+vYNlFXbY0/LqU+21ttmV+K471rPj0Jto7GPN4Zs6CaEr0g8COpDQNA6JoDv5Td0eIDWZ6c+ii5G9H+VjUCc6WprQIhepVkGzsJUjWlOrp66MeVwEElFdAk/PbXBvEUOWYTwi1uY6Y0trzMK31OvFOODKjWf6WHrf4tfgnAoGAOri6bX2D/zqpJT3mJ5MIVJJbn4D4Idx+TCUaVRSY1rBp+Y2ofW1W8ktu7xPO9/LwVQR7kJeosEBAFGTmGqll033ywu5+8X8J1bw6HCghkI0yHW752sOdfl30kXi3Ds8tQsvSEHRfnPb8yvWve2srZb9ubwOvpI0PtOIujZP4fYI="

  content, err := utils.DecryptCallback(platformPubKey, isvPriKey, callback)
  if nil != err {
	  //"decrypt failed"
  }
```
签名
```go
//utils.RsaSignBase64

// 此处为测试数据，正式使用是请使用真实的私钥
var priKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDBHBdHbQXsPT+EpAhLA9k2Q5O8GLCAUFLWYB57Uhc4ZNa2YUhjrTFvFZMFQuMjaVgdmFGTvqfGYUQBRldHFhf9kuXf5LPb+m0BJ/R5AWCyTcX7DHouoGODMfxkCZrimILwYWDkwhYTHr5hEV58nGRQtHOIVB5a4i/y4Z1vvX1MIjA+8OJ3zpaxXKkj+46OtfmjloUPGFSzz+rqrRRtMqYePLkWZ0J+CmIXM1Kwl/kgYUq/YGSYy9Q5vTojN9WBKzk7euOoCcsWtRrBQysdyM3yDPXjhnXx7G0nh07hSUh+rMDZ7Zst0lVQol/7kPXzNwh6eUmRGY9lfruIMS5kg1ydAgMBAAECggEAD4yQf0rTCEOiQq7mkAu+SLVGRwYB6EMPeH2C1tE0V3EfLM5GgugmK9ij3u+U1HweATwLjYbzgXDBhgzA6FNqGRvj8JQ8u0C92DL8Z2XqAFFs2JsXl3uIp761oOR5GTfIi0x7/c928ZEvKSe54PTCyxDMoLSNQSonTDpIb//k/+U4xEOQ1mjlSvlOM5ic7/kdw+G+aP/Hk/T6kg/vIblWQHx8SB3WYpLb/R6oPO+05X+zcQ+vVX1TrQ/amDp6/PouWjTF5hf48JEBdM8+xJzUwnalrG9U7pChfyGAOXQT1fbDdywBJXt6pZsT/mz1RkUC5Uto5/aVQGIDD+IPm/ZDbQKBgQDEwF09sjUb5hHEdmG28RmMf4E0JCOEzCvxiUpovobymqapLM5bf2oLNXqGenEAMbfFQatJFVKx6YBZwFIj/xzQJt8fL/jRzlbLijaANP+1JacvTsfXKBXS888FN3rkKisTlhmYXI+4EwA1wbcRkLDH3vezdVCi9cszQ9HvwkVfOwKBgQD7QvzD3pirXIJ64JizWTS4MJMko3CWepsq9UZ5uyoHWh7tSz86H/2y0FK10YpJEJeGtyPXlnU+uQwjYMJRPLlNv8180pjCJX2ZTW2drB2vOJvormhMhDIYAZtPAHu2dajzdy4VRuvFTtH4FpW/KjAJrTLK3ze3K95ACYVBJ8EmBwKBgB0K8DiNN724hmLjvqTMjiLpJ19U/lE5+jqbM3qmtTDWl0ddr9BdzH9/E2kKZefLbv8VJH2TQjO07hdRhk599/jZ5BGseSQvOyysaEMgj6ZjunwHOwSNjDspdiOk/uTzPIyVmY2eDDD1zRAiWi2jmBTI2vOIm7CSa75TgofLu4XFAoGBAJrFM4+vYNlFXbY0/LqU+21ttmV+K471rPj0Jto7GPN4Zs6CaEr0g8COpDQNA6JoDv5Td0eIDWZ6c+ii5G9H+VjUCc6WprQIhepVkGzsJUjWlOrp66MeVwEElFdAk/PbXBvEUOWYTwi1uY6Y0trzMK31OvFOODKjWf6WHrf4tfgnAoGAOri6bX2D/zqpJT3mJ5MIVJJbn4D4Idx+TCUaVRSY1rBp+Y2ofW1W8ktu7xPO9/LwVQR7kJeosEBAFGTmGqll033ywu5+8X8J1bw6HCghkI0yHW752sOdfl30kXi3Ds8tQsvSEHRfnPb8yvWve2srZb9ubwOvpI0PtOIujZP4fYI="
var content = "a=123！@#¥%……中文"
signature, error := utils.RsaSignBase64(content, priKey, crypto.SHA256)
if nil != error {
	//sign error
}
```
验签
```go
//utils.VerifySign

// 此处为测试数据
vvar pubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwRwXR20F7D0/hKQISwPZNkOTvBiwgFBS1mAee1IXOGTWtmFIY60xbxWTBULjI2lYHZhRk76nxmFEAUZXRxYX/ZLl3+Sz2/ptASf0eQFgsk3F+wx6LqBjgzH8ZAma4piC8GFg5MIWEx6+YRFefJxkULRziFQeWuIv8uGdb719TCIwPvDid86WsVypI/uOjrX5o5aFDxhUs8/q6q0UbTKmHjy5FmdCfgpiFzNSsJf5IGFKv2BkmMvUOb06IzfVgSs5O3rjqAnLFrUawUMrHcjN8gz144Z18extJ4dO4UlIfqzA2e2bLdJVUKJf+5D18zcIenlJkRmPZX67iDEuZINcnQIDAQAB"
var signature = "glTZg6lLl6oV4Ho15fAUegcVILlTwYJkbZO_Iz8AYUKTZ_1JP4AqAqSdm3GqjaukoNrDkxPGv2WW8plxYxtzsXjkzWiCMth5aShHgA7a9SXW0jfo365KPyVj0zFO2QIV9odHEnY1apwcAxvr54j4d5SHoC3vKUczZ20txTsNjcG9ifi1AoJhblILxKL2NO0tdIzTMQCRaBdOXUOdnL7RgP1qPew5yJT4e1QdtTjkirCKJurm4SumOA3Uroz-G-9MUZgiTkU4RXrEvu-rJPlqfJPsITYoWLsuPy1Gfne_5j-IgChXpoHacI0s-NlzKmyjsFt3-5aUYDd0cFw58ErUXw"
var data = "{\"result\":{\"requestId\":\"requestId\",\"errorMsg\":\"exception.record.not.found.transferDomesticOrder|merchantId:[null],requestId:[requestId]\",\"status\":\"FAILED\"}}"
if !utils.VerifySign(data, signature, pubKey, crypto.SHA256) {
     //verify failed
}
```
