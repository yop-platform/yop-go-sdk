# YOP Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/yop-platform/yop-go-sdk.svg)](https://pkg.go.dev/github.com/yop-platform/yop-go-sdk)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[阅读中文文档](README_zh-CN.md)

A Go SDK specifically designed for seamless interaction with YOP (YeePay Open Platform) APIs.

## 📋 Overview

This SDK provides a convenient way to integrate YeePay payment and other services into your Go applications. It handles request signing, signature verification, and API communication, allowing you to focus on your application logic.

**Key Features:**

- **Secure and Reliable**: Implements YOP API's RSA signature requirements, supports UTF-8 encoding for international characters
- **Simple to Use**: Provides a clean API interface, reducing integration complexity
- **Comprehensive**: Supports all YOP Open Platform interfaces, including payment, refund, query, and more

**Use Cases:**

1. Request YOP Open Platform interfaces
2. Decrypt YOP Open Platform callback content
3. Construct YOP Cashier signature

## 📥 Installation

Install this package using Go modules:

```bash
go get github.com/yop-platform/yop-go-sdk
```

## ⚙️ Configuration

YOP Go SDK configuration is primarily set through the `YopRequest` object. Here are the main configuration items:

### Basic Configuration

- **AppId**: Your YeePay application AppKey, provided by YeePay
- **IsvPriKey**: Your application's private key, used for signing requests
- **Timeout**: Request timeout, default is 10 seconds

### Certificate Types

The SDK supports the following certificate types:

```go
// Defined in the request package
const (
    RSA2048 = "RSA2048" // RSA2048 algorithm
)
```

### Configuration Example

```go
var priKey = &request.IsvPriKey{Value: "Your private key content", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/api/path")
yopRequest.AppId = "Your AppId"
yopRequest.IsvPriKey = priKey
yopRequest.Timeout = 15 // Set timeout to 15 seconds (optional)
```

## 🚀 Usage / Quick Start

### Import Packages

```go
import (
    "github.com/yop-platform/yop-go-sdk/yop/client"
    "github.com/yop-platform/yop-go-sdk/yop/constants"
    "github.com/yop-platform/yop-go-sdk/yop/request"
    "github.com/yop-platform/yop-go-sdk/yop/utils"
)
```

### GET Request Example

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
//yopResp.Result is the request result
```

### POST Form Request Example

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
//yopResp.Result is the request result
```

### POST JSON Request Example

```go
var priKey = request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
var yopRequest = request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
// Set JSON request payload
var params = map[string]any{}
params["merchantId"] = "1595815987915711"
params["requestId"] = "requestId"
result.Content = utils.ParseToJsonStr(params)

yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{ 
    // request failed
}
//yopResp.Result is the request result
```

### File Upload Request Example

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
// yopResp.Result is the upload request result
```

### File Download Request Example

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
// yopResp.Content is the file content
```

### Callback Processing Example

```go
//utils.DecryptCallback

var callback = "Ars6jASSiylO70_VJDQ5SFU1zQwaI36kG5WhlSKHjkGdU3fEVEkkbhvAxKjOTUiw9vF7RMnmGKQQWAuV8jCKaOpMNjIEMHehBaPASwTiEE946CcbOeoNILGHf0o20xj2gqqvkQToFXEMNiic7bcYbfi0PxIrR6loBZnW-m5bqzB5RXLibiSjGlmr5CDnxV4tZXmYlkkeN2BcT4msWjfCtuaTMK_fN77WJcCMlW7ffqiN5yIOeqB4QBb5lOnClTRW4DThKPOMkXupAM2AnPxTkDp4n9lh-SK56zLuafk1bQhWUNcS9L4YEKZGJIjP7DY20TAWEr3yXo8w0w0VtB13Ig$Xf6fETKWcLTudBh2HluGSQTqhBRJa6EXHhXlMryWW8Y384RjVwIfpQm19RmTgkoqRc2tNcTWxRIW6itIS62DrzixlqRa099jx21uGqt8FCpvdWwnwlC16SgkeU_5NnrpjA_WQ0XW9RhNxzuQmwfxHGbtnth4vNXWswcSm23j3KQaXFjVP5Ws1uYVCxYSLMxqJE7a56DNWONGcGJJsc0KTCc7cdfr8n24emAaPCNteIG2RM8F17pRxY5yVnguTSZPXmhBlyI25xS7rciWzKZLp2Kfh_JCivABbA-_5Vf3VWPmjITs-TR5HlGVFbnT0eOUMUepXUemjjP8R0f8cBeH2NKej6QjQL99tvlrrxg_QfmezE0WTCITCNDBhpbHiq90lFyLjwlWNDTRo8rhjouSlMA9Ae_b-B4eZorDRVxw3BWywdyo2FzNk-dUDeBVaIth9YsaMGsq9XivGjlnnx3YEVfEtuVSvEm1xBdYsTHcM02nMwZb8Ze2WL1kIFo8IFM0$AES$SHA256"
// This is test data, please use the real platform public key for production
var platformPubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7LqdMV7ZeOWUwVp0duSucTr4VwUNHtYLlWEUWlBtDQDEPhx0WZZdw2DxEbQqMQM5BjXZACYlhEdPt0HicDthOIUeUt8JNcvgq06vIE958RzgVBa5z3zvMLYWJIZaUyxsxC7Us06eNiB+du0rEBxUckru41ZSu/DX9jssFC+l5459b3WWELNf2fXqJyfb4f8GuGk8enXgJdxBUcmwgaEQxJjWkPqhzSiRy9GKjcXBdCkzCYR4xmLkHe6K0YFiBxax7lOni3zVOsvHC9XdhbepwB9fMkHbZXS/LJf5aS5ltendObpVrAD9kck7bIQzsrM49/SG/dYmbtm139I6ygsCzQIDAQAB"
var isvPriKey = "<Private Key Text>"

content, err := utils.DecryptCallback(platformPubKey, isvPriKey, callback)
if nil != err {
    //"decrypt failed"
}
```

### Signature Operation Example

```go
//utils.RsaSignBase64

// This is test data, please use the real private key for production
var priKey = "<Private Key Text>"
var content = "a=123！@#¥%……Chinese"
signature, error := utils.RsaSignBase64(content, priKey, crypto.SHA256)
if nil != error {
    //sign error
}
```

### Signature Verification Example

```go
//utils.VerifySign

// This is test data
var pubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwRwXR20F7D0/hKQISwPZNkOTvBiwgFBS1mAee1IXOGTWtmFIY60xbxWTBULjI2lYHZhRk76nxmFEAUZXRxYX/ZLl3+Sz2/ptASf0eQFgsk3F+wx6LqBjgzH8ZAma4piC8GFg5MIWEx6+YRFefJxkULRziFQeWuIv8uGdb719TCIwPvDid86WsVypI/uOjrX5o5aFDxhUs8/q6q0UbTKmHjy5FmdCfgpiFzNSsJf5IGFKv2BkmMvUOb06IzfVgSs5O3rjqAnLFrUawUMrHcjN8gz144Z18extJ4dO4UlIfqzA2e2bLdJVUKJf+5D18zcIenlJkRmPZX67iDEuZINcnQIDAQAB"
var signature = "glTZg6lLl6oV4Ho15fAUegcVILlTwYJkbZO_Iz8AYUKTZ_1JP4AqAqSdm3GqjaukoNrDkxPGv2WW8plxYxtzsXjkzWiCMth5aShHgA7a9SXW0jfo365KPyVj0zFO2QIV9odHEnY1apwcAxvr54j4d5SHoC3vKUczZ20txTsNjcG9ifi1AoJhblILxKL2NO0tdIzTMQCRaBdOXUOdnL7RgP1qPew5yJT4e1QdtTjkirCKJurm4SumOA3Uroz-G-9MUZgiTkU4RXrEvu-rJPlqfJPsITYoWLsuPy1Gfne_5j-IgChXpoHacI0s-NlzKmyjsFt3-5aUYDd0cFw58ErUXw"
var data = "{\"result\":{\"requestId\":\"requestId\",\"errorMsg\":\"exception.record.not.found.transferDomesticOrder|merchantId:[null],requestId:[requestId]\",\"status\":\"FAILED\"}}"
if !utils.VerifySign(data, signature, pubKey, crypto.SHA256) {
     //verify failed
}
```

## 📖 API Reference

### YopRequest

`YopRequest` is the main object for initiating YOP API requests.

#### Main Properties

- `AppId`: Application identifier
- `IsvPriKey`: Application private key
- `Timeout`: Request timeout (seconds)
- `Content`: Request content (for JSON requests)

#### Main Methods

- `NewYopRequest(httpMethod, apiUri string)`: Create a new request object
- `AddParam(paramName, paramValue string)`: Add request parameter
- `AddFile(paramName string, file *os.File)`: Add upload file

### YopResponse

`YopResponse` is the encapsulation of YOP API responses.

#### Main Properties

- `State`: Response status
- `Result`: Request result
- `Error`: Error information
- `Content`: Response content (for file downloads)

### Utility Functions

#### Callback Processing

- `utils.DecryptCallback(platformPubKey, isvPriKey, callback string) (string, error)`: Decrypt callback content

#### Signature Operations

- `utils.RsaSignBase64(content, priKey string, hash crypto.Hash) (string, error)`: Generate signature
- `utils.VerifySign(data, signature, pubKey string, hash crypto.Hash) bool`: Verify signature

## 🤝 Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please submit an issue or pull request.

---

## 📚 Developer Resources

*   **Official Java SDK(RSA) Usage Guide**: [https://open.yeepay.com/docs/platform/sdk_guide/java-sdk-guide](https://open.yeepay.com/docs/platform/sdk_guide/java-sdk-guide)
*   **API Interface List and Documentation**: [https://open.yeepay.com/docs-v2](https://open.yeepay.com/docs/api-list)
*   **Encryption Machine Integration Guide**: [https://open.yeepay.com/docs/open/platform-doc/sdk_guide-sm/encryptor-support](https://open.yeepay.com/docs/open/platform-doc/sdk_guide-sm/encryptor-support)

---

## 📜 License

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyop-platform%2Fyop-go-sdk.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyop-platform%2Fyop-go-sdk?ref=badge_large)

This SDK follows the [Apache License 2.0](LICENSE) open source license agreement.

---

We are committed to providing an excellent developer experience. If you encounter any issues or have any suggestions during use, please feel free to contact us through official channels.
