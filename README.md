# yop-go-sdk
Yeepay openapi SDK fo Go
## 本项目的使用场景
1.请求易宝开放平台接口 </br>
2.解密易宝开放平台回调内容</br>
3.构造易宝收银台签名
## 使用
1.引入
```go
import yopGoSdk "github.com/yop-platform/yop-go-sdk"
```
2.调用接口</br>
Get请求
```go

var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = &request.YopRequest{}
yopRequest.AppId = "appId"
yopRequest.ApiUri = "/rest/v1.0/test-wdc/product-query/query-for-doc"
yopRequest.HttpMethod = constants.GET_HTTP_METHOD
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := DefaultClient.Request(yopRequest)
```
Post Form请求
```go

var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = &request.YopRequest{}
yopRequest.AppId = "appId"
yopRequest.ApiUri = "/rest/v1.0/test-wdc/product-query/query-for-doc"
yopRequest.HttpMethod = constants.POST_HTTP_METHOD
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := DefaultClient.Request(yopRequest)
```
Post Json请求
```go

var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = &request.YopRequest{}
yopRequest.AppId = "appId"
yopRequest.ApiUri = "/rest/v1.0/test-wdc/product-query/query-for-doc"
yopRequest.HttpMethod = constants.POST_HTTP_METHOD
yopRequest.IsvPriKey = priKey
// 设置json请求报文
yopRequest.Content = "{\"merchantId\":\"1595815987915711\",\"requestId\":\"requestId\"}"
yopResp, err := DefaultClient.Request(yopRequest)
```

文件上传请求
```go

var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = &request.YopRequest{}
yopRequest.AppId = "appId"
yopRequest.ApiUri = "/rest/v1.0/test-wdc/product-query/query-for-doc"
yopRequest.ServerRoot = "http://ycetest.yeepay.com:30228/yop-center"
yopRequest.HttpMethod = constants.POST_HTTP_METHOD
yopRequest.IsvPriKey = priKey
result.AddFile("file", f)
yopResp, err := DefaultClient.Request(yopRequest)
```

文件下载请求
```go

var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = &request.YopRequest{}
yopRequest.AppId = "appId"
yopRequest.ApiUri = "/rest/v1.0/test-wdc/product-query/query-for-doc"
yopRequest.HttpMethod = constants.GET_HTTP_METHOD
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := DefaultClient.Request(yopRequest)
```