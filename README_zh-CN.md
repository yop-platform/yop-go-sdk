# YOP Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/yop-platform/yop-go-sdk.svg)](https://pkg.go.dev/github.com/yop-platform/yop-go-sdk)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/yop-platform/yop-typescript-sdk)

一个专为与 YOP（易宝开放平台）API 进行无缝交互而设计的 Go SDK。

## 📋 概述

此 SDK 提供了一种便捷的方式，可将易宝支付及其他服务集成到您的 Go 应用中。它负责处理请求签名、签名验证和 API 通信，让您可以专注于您的应用逻辑。

**主要特性：**

- **安全可靠**：实现 YOP API 的 RSA 签名要求，支持 UTF-8 编码处理国际字符
- **简单易用**：提供简洁的 API 接口，降低集成难度
- **功能完整**：支持所有 YOP 开放平台接口，包括POST、GET、文件上传、文件下载等功能

**使用场景：**

1. 请求易宝开放平台接口
2. 解密易宝开放平台回调内容
3. 构造易宝收银台签名

## 📥 安装

使用 Go 模块安装此包：

```bash
go get github.com/yop-platform/yop-go-sdk
```

## ⚙️ 配置

YOP Go SDK 的配置主要通过 `YopRequest` 对象进行设置。以下是主要的配置项：

### 基本配置

- **AppId**: 您的易宝应用 AppKey，由易宝提供
- **IsvPriKey**: 您应用的私钥，用于签名请求
- **Timeout**: 请求超时时间，默认为 10 秒

### 证书类型

SDK 支持以下证书类型：

```go
// 在 request 包中定义
const (
    RSA2048 = "RSA2048" // RSA2048 算法
)
```

### 配置示例

```go
var priKey = &request.IsvPriKey{Value: "您的私钥内容", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/api/path")
yopRequest.AppId = "您的AppId"
yopRequest.IsvPriKey = priKey
yopRequest.Timeout = 15 // 设置超时时间为 15 秒（可选）
```

## 🚀 用法 / 快速开始

### 引入包

```go
import (
    "github.com/yop-platform/yop-go-sdk/yop/client"
    "github.com/yop-platform/yop-go-sdk/yop/constants"
    "github.com/yop-platform/yop-go-sdk/yop/request"
    "github.com/yop-platform/yop-go-sdk/yop/utils"
)
```

### GET 请求示例

```go
var priKey = &request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.GET_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{
    // request failed
}
//yopResp.Result为请求结果
```

### POST Form 请求示例

```go
var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
    // request failed
}
//yopResp.Result为请求结果
```

### POST JSON 请求示例

```go
var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
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

### 文件上传请求示例

```go
var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
result.AddFile("file", f)
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
    // request failed
}
// yopResp.Result为上传请求结果
```

### 文件下载请求示例

```go
var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.GET_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
    // request failed
}
// yopResp.Content为文件内容
```

### 回调处理示例

```go
//utils.DecryptCallback

var callback = "Ars6jASSiylO70_VJDQ5SFU1zQwaI36kG5WhlSKHjkGdU3fEVEkkbhvAxKjOTUiw9vF7RMnmGKQQWAuV8jCKaOpMNjIEMHehBaPASwTiEE946CcbOeoNILGHf0o20xj2gqqvkQToFXEMNiic7bcYbfi0PxIrR6loBZnW-m5bqzB5RXLibiSjGlmr5CDnxV4tZXmYlkkeN2BcT4msWjfCtuaTMK_fN77WJcCMlW7ffqiN5yIOeqB4QBb5lOnClTRW4DThKPOMkXupAM2AnPxTkDp4n9lh-SK56zLuafk1bQhWUNcS9L4YEKZGJIjP7DY20TAWEr3yXo8w0w0VtB13Ig$Xf6fETKWcLTudBh2HluGSQTqhBRJa6EXHhXlMryWW8Y384RjVwIfpQm19RmTgkoqRc2tNcTWxRIW6itIS62DrzixlqRa099jx21uGqt8FCpvdWwnwlC16SgkeU_5NnrpjA_WQ0XW9RhNxzuQmwfxHGbtnth4vNXWswcSm23j3KQaXFjVP5Ws1uYVCxYSLMxqJE7a56DNWONGcGJJsc0KTCc7cdfr8n24emAaPCNteIG2RM8F17pRxY5yVnguTSZPXmhBlyI25xS7rciWzKZLp2Kfh_JCivABbA-_5Vf3VWPmjITs-TR5HlGVFbnT0eOUMUepXUemjjP8R0f8cBeH2NKej6QjQL99tvlrrxg_QfmezE0WTCITCNDBhpbHiq90lFyLjwlWNDTRo8rhjouSlMA9Ae_b-B4eZorDRVxw3BWywdyo2FzNk-dUDeBVaIth9YsaMGsq9XivGjlnnx3YEVfEtuVSvEm1xBdYsTHcM02nMwZb8Ze2WL1kIFo8IFM0$AES$SHA256"
// 此处为测试数据，正式使用是请使用真实的平台公钥
var platformPubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7LqdMV7ZeOWUwVp0duSucTr4VwUNHtYLlWEUWlBtDQDEPhx0WZZdw2DxEbQqMQM5BjXZACYlhEdPt0HicDthOIUeUt8JNcvgq06vIE958RzgVBa5z3zvMLYWJIZaUyxsxC7Us06eNiB+du0rEBxUckru41ZSu/DX9jssFC+l5459b3WWELNf2fXqJyfb4f8GuGk8enXgJdxBUcmwgaEQxJjWkPqhzSiRy9GKjcXBdCkzCYR4xmLkHe6K0YFiBxax7lOni3zVOsvHC9XdhbepwB9fMkHbZXS/LJf5aS5ltendObpVrAD9kck7bIQzsrM49/SG/dYmbtm139I6ygsCzQIDAQAB"
var isvPriKey = "<私钥文本>"

content, err := utils.DecryptCallback(platformPubKey, isvPriKey, callback)
if nil != err {
    //"decrypt failed"
}
```

### 签名操作示例

```go
//utils.RsaSignBase64

// 此处为测试数据，正式使用是请使用真实的私钥
var priKey = "<私钥文本>"
var content = "a=123！@#¥%……中文"
signature, error := utils.RsaSignBase64(content, priKey, crypto.SHA256)
if nil != error {
    //sign error
}
```

### 验证签名示例

```go
//utils.VerifySign

// 此处为测试数据
var pubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwRwXR20F7D0/hKQISwPZNkOTvBiwgFBS1mAee1IXOGTWtmFIY60xbxWTBULjI2lYHZhRk76nxmFEAUZXRxYX/ZLl3+Sz2/ptASf0eQFgsk3F+wx6LqBjgzH8ZAma4piC8GFg5MIWEx6+YRFefJxkULRziFQeWuIv8uGdb719TCIwPvDid86WsVypI/uOjrX5o5aFDxhUs8/q6q0UbTKmHjy5FmdCfgpiFzNSsJf5IGFKv2BkmMvUOb06IzfVgSs5O3rjqAnLFrUawUMrHcjN8gz144Z18extJ4dO4UlIfqzA2e2bLdJVUKJf+5D18zcIenlJkRmPZX67iDEuZINcnQIDAQAB"
var signature = "glTZg6lLl6oV4Ho15fAUegcVILlTwYJkbZO_Iz8AYUKTZ_1JP4AqAqSdm3GqjaukoNrDkxPGv2WW8plxYxtzsXjkzWiCMth5aShHgA7a9SXW0jfo365KPyVj0zFO2QIV9odHEnY1apwcAxvr54j4d5SHoC3vKUczZ20txTsNjcG9ifi1AoJhblILxKL2NO0tdIzTMQCRaBdOXUOdnL7RgP1qPew5yJT4e1QdtTjkirCKJurm4SumOA3Uroz-G-9MUZgiTkU4RXrEvu-rJPlqfJPsITYoWLsuPy1Gfne_5j-IgChXpoHacI0s-NlzKmyjsFt3-5aUYDd0cFw58ErUXw"
var data = "{\"result\":{\"requestId\":\"requestId\",\"errorMsg\":\"exception.record.not.found.transferDomesticOrder|merchantId:[null],requestId:[requestId]\",\"status\":\"FAILED\"}}"
if !utils.VerifySign(data, signature, pubKey, crypto.SHA256) {
     //verify failed
}
```

## 📖 API 参考

### YopRequest

`YopRequest` 是发起 YOP API 请求的主要对象。

#### 主要属性

- `AppId`: 应用标识
- `IsvPriKey`: 应用私钥
- `Timeout`: 请求超时时间（秒）
- `Content`: 请求内容（用于 JSON 请求）

#### 主要方法

- `NewYopRequest(httpMethod, apiUri string)`: 创建新的请求对象
- `AddParam(paramName, paramValue string)`: 添加请求参数
- `AddFile(paramName string, file *os.File)`: 添加上传文件

### YopResponse

`YopResponse` 是 YOP API 响应的封装。

#### 主要属性

- `State`: 响应状态
- `Result`: 请求结果
- `Error`: 错误信息
- `Content`: 响应内容（用于文件下载）

### 工具函数

#### 回调处理

- `utils.DecryptCallback(platformPubKey, isvPriKey, callback string) (string, error)`: 解密回调内容

#### 签名操作

- `utils.RsaSignBase64(content, priKey string, hash crypto.Hash) (string, error)`: 生成签名
- `utils.VerifySign(data, signature, pubKey string, hash crypto.Hash) bool`: 验证签名

## 🤝 贡献

欢迎贡献！如果您发现任何问题或有改进建议，请提交 issue 或 pull request。

---

## 📚 开发者资源

*   **官方Java SDK(RSA)使用指南**: [https://open.yeepay.com/docs/platform/sdk_guide/java-sdk-guide](https://open.yeepay.com/docs/platform/sdk_guide/java-sdk-guide)
*   **API接口列表与文档**: [https://open.yeepay.com/docs-v2](https://open.yeepay.com/docs/api-list)
*   **加密机对接指南**: [https://open.yeepay.com/docs/open/platform-doc/sdk_guide-sm/encryptor-support](https://open.yeepay.com/docs/open/platform-doc/sdk_guide-sm/encryptor-support)

---

## 📜 License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyop-platform%2Fyop-go-sdk.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyop-platform%2Fyop-go-sdk?ref=badge_large)

本 SDK 遵循 [Apache License 2.0](LICENSE) 开源许可协议。

---

我们致力于提供卓越的开发者体验。如果您在使用过程中遇到任何问题或有任何建议，欢迎通过官方渠道与我们联系。
