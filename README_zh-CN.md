# YOP Go SDK

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/yop-platform/yop-go-sdk.svg)](https://pkg.go.dev/github.com/yop-platform/yop-go-sdk)
[![CI/CD](https://github.com/yop-platform/yop-go-sdk/workflows/CI%2FCD/badge.svg)](https://github.com/yop-platform/yop-go-sdk/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/yop-platform/yop-go-sdk)](https://goreportcard.com/report/github.com/yop-platform/yop-go-sdk)
[![codecov](https://codecov.io/gh/yop-platform/yop-go-sdk/branch/main/graph/badge.svg)](https://codecov.io/gh/yop-platform/yop-go-sdk)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GitHub release](https://img.shields.io/github/release/yop-platform/yop-go-sdk.svg)](https://github.com/yop-platform/yop-go-sdk/releases)
[![Go version](https://img.shields.io/github/go-mod/go-version/yop-platform/yop-go-sdk)](https://github.com/yop-platform/yop-go-sdk)
[![Ask DeepWiki](https://deepwiki.com/badge.svg)](https://deepwiki.com/yop-platform/yop-typescript-sdk)

[English](README.md) | 中文

</div>

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
# 安装最新版本
go get github.com/yop-platform/yop-go-sdk

# 安装指定版本
go get github.com/yop-platform/yop-go-sdk@v1.4.40
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
priKey := request.IsvPriKey{Value: "您的私钥内容", CertType: request.RSA2048}
yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/api/path")
yopRequest.AppId = "您的AppId"
yopRequest.IsvPriKey = priKey
yopRequest.Timeout = 15 * time.Second // 设置超时时间为 15 秒（可选）
```
## 🚀 使用方法 / 快速开始

### 导入包
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
priKey := request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
yopRequest := request.NewYopRequest(constants.GET_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{
    // 请求失败
}
// yopResp.Result 是请求结果
```
### POST Form 请求示例

```go
priKey := request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddParam("paramName", "paramValue")
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{
    // 请求失败
}
// yopResp.Result 是请求结果
```
### POST JSON 请求示例

```go
priKey := request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
// 设置 JSON 请求载荷
params := map[string]any{}
params["merchantId"] = "1595815987915711"
params["requestId"] = "requestId"
yopRequest.Content = utils.ParseToJsonStr(params)

yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{
    // 请求失败
}
// yopResp.Result 是请求结果
```
### 文件上传请求示例

```go
priKey := request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
yopRequest := request.NewYopRequest(constants.POST_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")
yopRequest.AppId = "appId"
yopRequest.IsvPriKey = priKey
yopRequest.AddFile("file", f)
yopResp, err := client.DefaultClient.Request(yopRequest)
if nil != err{
    // 请求失败
}
// yopResp.Result 是上传请求结果
```

### 文件下载请求示例

```go
priKey := request.IsvPriKey{Value: "isvPriKey", CertType: request.RSA2048}
yopRequest := request.NewYopRequest(constants.GET_HTTP_METHOD, "/rest/v1.0/test/product-query/query-for-doc")

var callback = "Ars6jASSiylO70_VJDQ5SFU1zQwaI36kG5WhlSKHjkGdU3fEVEkkbhvAxKjOTUiw9vF7RMnmGKQQWAuV8jCKaOpMNjIEMHehBaPASwTiEE946CcbOeoNILGHf0o20xj2gqqvkQToFXEMNiic7bcYbfi0PxIrR6loBZnW-m5bqzB5RXLibiSjGlmr5CDnxV4tZXmYlkkeN2BcT4msWjfCtuaTMK_fN77WJcCMlW7ffqiN5yIOeqB4QBb5lOnClTRW4DThKPOMkXupAM2AnPxTkDp4n9lh-SK56zLuafk1bQhWUNcS9L4YEKZGJIjP7DY20TAWEr3yXo8w0w0VtB13Ig$Xf6fETKWcLTudBh2HluGSQTqhBRJa6EXHhXlMryWW8Y384RjVwIfpQm19RmTgkoqRc2tNcTWxRIW6itIS62DrzixlqRa099jx21uGqt8FCpvdWwnwlC16SgkeU_5NnrpjA_WQ0XW9RhNxzuQmwfxHGbtnth4vNXWswcSm23j3KQaXFjVP5Ws1uYVCxYSLMxqJE7a56DNWONGcGJJsc0KTCc7cdfr8n24emAaPCNteIG2RM8F17pRxY5yVnguTSZPXmhBlyI25xS7rciWzKZLp2Kfh_JCivABbA-_5Vf3VWPmjITs-TR5HlGVFbnT0eOUMUepXUemjjP8R0f8cBeH2NKej6QjQL99tvlrrxg_QfmezE0WTCITCNDBhpbHiq90lFyLjwlWNDTRo8rhjouSlMA9Ae_b-B4eZorDRVxw3BWywdyo2FzNk-dUDeBVaIth9YsaMGsq9XivGjlnnx3YEVfEtuVSvEm1xBdYsTHcM02nMwZb8Ze2WL1kIFo8IFM0$AES$SHA256"
// 此处为测试数据，正式使用是请使用真实的平台公钥
var platformPubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA7LqdMV7ZeOWUwVp0duSucTr4VwUNHtYLlWEUWlBtDQDEPhx0WZZdw2DxEbQqMQM5BjXZACYlhEdPt0HicDthOIUeUt8JNcvgq06vIE958RzgVBa5z3zvMLYWJIZaUyxsxC7Us06eNiB+du0rEBxUckru41ZSu/DX9jssFC+l5459b3WWELNf2fXqJyfb4f8GuGk8enXgJdxBUcmwgaEQxJjWkPqhzSiRy9GKjcXBdCkzCYR4xmLkHe6K0YFiBxax7lOni3zVOsvHC9XdhbepwB9fMkHbZXS/LJf5aS5ltendObpVrAD9kck7bIQzsrM49/SG/dYmbtm139I6ygsCzQIDAQAB"
var isvPriKey = "<私钥文本>"

content, err := utils.DecryptCallback(platformPubKey, isvPriKey, callback)
if nil != err {
    // 解密失败
}
```

### 签名操作示例

```go
// utils.RsaSignBase64

// 此处为测试数据，正式使用时请使用真实的私钥
priKey := "<私钥文本>"
content := "a=123！@#¥%……中文"
signature, err := utils.RsaSignBase64(content, priKey, crypto.SHA256)
if nil != err {
    // 签名错误
}
```

### 验证签名示例

```go
// utils.VerifySign

// 此处为测试数据
pubKey := "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwRwXR20F7D0/hKQISwPZNkOTvBiwgFBS1mAee1IXOGTWtmFIY60xbxWTBULjI2lYHZhRk76nxmFEAUZXRxYX/ZLl3+Sz2/ptASf0eQFgsk3F+wx6LqBjgzH8ZAma4piC8GFg5MIWEx6+YRFefJxkULRziFQeWuIv8uGdb719TCIwPvDid86WsVypI/uOjrX5o5aFDxhUs8/q6q0UbTKmHjy5FmdCfgpiFzNSsJf5IGFKv2BkmMvUOb06IzfVgSs5O3rjqAnLFrUawUMrHcjN8gz144Z18extJ4dO4UlIfqzA2e2bLdJVUKJf+5D18zcIenlJkRmPZX67iDEuZINcnQIDAQAB"
signature := "glTZg6lLl6oV4Ho15fAUegcVILlTwYJkbZO_Iz8AYUKTZ_1JP4AqAqSdm3GqjaukoNrDkxPGv2WW8plxYxtzsXjkzWiCMth5aShHgA7a9SXW0jfo365KPyVj0zFO2QIV9odHEnY1apwcAxvr54j4d5SHoC3vKUczZ20txTsNjcG9ifi1AoJhblILxKL2NO0tdIzTMQCRaBdOXUOdnL7RgP1qPew5yJT4e1QdtTjkirCKJurm4SumOA3Uroz-G-9MUZgiTkU4RXrEvu-rJPlqfJPsITYoWLsuPy1Gfne_5j-IgChXpoHacI0s-NlzKmyjsFt3-5aUYDd0cFw58ErUXw"
data := "{\"result\":{\"requestId\":\"requestId\",\"errorMsg\":\"exception.record.not.found.transferDomesticOrder|merchantId:[null],requestId:[requestId]\",\"status\":\"FAILED\"}}"
if !utils.VerifySign(data, signature, pubKey, crypto.SHA256) {
     // 验证失败
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

## 🔧 高级配置

### 环境配置

SDK 支持多种环境配置：

```go
// 生产环境（默认）
yopRequest.ServerRoot = "https://openapi.yeepay.com/yop-center"

// 测试环境
yopRequest.ServerRoot = "https://ycetest.yeepay.com:30228/yop-center"

// YOS 文件服务
yopRequest.ServerRoot = "https://yos.yeepay.com/yop-center"
```

### 自定义HTTP客户端

```go
import (
    "net/http"
    "time"
)

// 创建自定义HTTP客户端
customClient := &http.Client{
    Timeout: 30 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
        IdleConnTimeout:     90 * time.Second,
    },
}

// 使用自定义客户端
yopClient := client.YopClient{Client: customClient}
yopResp, err := yopClient.Request(yopRequest)
```

### 日志配置

SDK 提供了统一的日志系统，可以轻松配置：

```go
import (
    "os"
    "github.com/yop-platform/yop-go-sdk/yop/utils"
    "github.com/sirupsen/logrus"
)

// 基础日志配置
// 设置日志级别 (Debug, Info, Warn, Error)
utils.SetLogLevel(logrus.InfoLevel)

// 完全禁用日志
utils.DisableLogging()

// 重新启用日志
utils.EnableLogging()

// 自定义日志器配置
customLogger := logrus.New()
customLogger.SetOutput(os.Stdout)
customLogger.SetLevel(logrus.DebugLevel)
customLogger.SetFormatter(&logrus.JSONFormatter{})
utils.SetLogger(customLogger)

// 设置自定义格式化器
utils.SetLogFormatter(&logrus.JSONFormatter{
    TimestampFormat: "2006-01-02 15:04:05",
})

// 示例：生产环境日志设置
func setupProductionLogging() {
    // 生产环境使用结构化JSON日志
    utils.SetLogFormatter(&logrus.JSONFormatter{
        TimestampFormat: "2006-01-02T15:04:05.000Z",
    })

    // 设置合适的日志级别
    utils.SetLogLevel(logrus.WarnLevel)

    // 可选：日志输出到文件而不是标准输出
    logFile, err := os.OpenFile("yop-sdk.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err == nil {
        customLogger := logrus.New()
        customLogger.SetOutput(logFile)
        customLogger.SetLevel(logrus.WarnLevel)
        customLogger.SetFormatter(&logrus.JSONFormatter{})
        utils.SetLogger(customLogger)
    }
}
```

#### 可用的日志级别

- `logrus.DebugLevel`: 详细的调试信息
- `logrus.InfoLevel`: 一般信息（默认）
- `logrus.WarnLevel`: 警告消息
- `logrus.ErrorLevel`: 仅错误消息

#### 日志输出示例

```go
// SDK 会自动记录请求信息：
// time="2023-12-01T10:30:45Z" level=info msg="requestId:abc123-def456"
// time="2023-12-01T10:30:45Z" level=info msg="authString:yop-auth-v3/..."
// time="2023-12-01T10:30:45Z" level=info msg="statusCode:200"
```

## 🚨 错误处理

### 常见错误类型

```go
yopResp, err := client.DefaultClient.Request(yopRequest)
if err != nil {
    // 网络错误或请求构建错误
    log.Printf("请求失败: %v", err)
    return
}

// 检查业务错误
if yopResp.Result != nil {
    result := yopResp.Result.(map[string]interface{})
    if status, ok := result["status"]; ok && status != "SUCCESS" {
        log.Printf("业务错误: %v", result["errorMsg"])
        return
    }
}
```

### 重试机制

```go
func requestWithRetry(yopRequest *request.YopRequest, maxRetries int) (*response.YopResponse, error) {
    var lastErr error

    for i := 0; i <= maxRetries; i++ {
        yopResp, err := client.DefaultClient.Request(yopRequest)
        if err == nil {
            return yopResp, nil
        }

        lastErr = err
        if i < maxRetries {
            time.Sleep(time.Duration(i+1) * time.Second) // 指数退避
        }
    }

    return nil, fmt.Errorf("请求失败，已重试 %d 次: %v", maxRetries, lastErr)
}
```

## 📊 性能优化

### 连接池配置

```go
// 优化HTTP传输配置
transport := &http.Transport{
    MaxIdleConns:        100,
    MaxIdleConnsPerHost: 10,
    IdleConnTimeout:     90 * time.Second,
    TLSHandshakeTimeout: 10 * time.Second,
}

customClient := &http.Client{
    Transport: transport,
    Timeout:   30 * time.Second,
}
```

### 批量请求处理

```go
func processBatchRequests(requests []*request.YopRequest) {
    const maxConcurrency = 10
    semaphore := make(chan struct{}, maxConcurrency)
    var wg sync.WaitGroup

    for _, req := range requests {
        wg.Add(1)
        go func(r *request.YopRequest) {
            defer wg.Done()
            semaphore <- struct{}{} // 获取信号量
            defer func() { <-semaphore }() // 释放信号量

            resp, err := client.DefaultClient.Request(r)
            if err != nil {
                log.Printf("请求失败: %v", err)
                return
            }
            // 处理响应...
        }(req)
    }

    wg.Wait()
}
```

## 🧪 测试

### 单元测试

```bash
# 运行所有测试
go test ./...

# 运行测试并显示覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### 基准测试

```bash
# 运行基准测试
go test -bench=. ./...

# 运行基准测试并显示内存分配
go test -bench=. -benchmem ./...
```

## 🤝 贡献指南

我们欢迎所有形式的贡献！在贡献之前，请阅读以下指南：

### 开发环境设置

1. **克隆仓库**
   ```bash
   git clone https://github.com/yop-platform/yop-go-sdk.git
   cd yop-go-sdk
   ```

2. **安装依赖**
   ```bash
   go mod download
   ```

3. **运行测试**
   ```bash
   go test ./...
   ```

4. **代码格式化**
   ```bash
   go fmt ./...
   goimports -w .
   ```

### 提交规范

- 使用清晰的提交信息
- 遵循 [Conventional Commits](https://www.conventionalcommits.org/) 规范
- 确保所有测试通过
- 添加必要的测试用例

### Pull Request 流程

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

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
