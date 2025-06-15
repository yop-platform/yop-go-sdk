// Package request
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/16 3:22 PM
package request

import (
	"encoding/json"
	"fmt"
	"html/template"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/yop-platform/yop-go-sdk/yop/utils"
)

const (
	SERVER_ROOT             = "https://openapi.yeepay.com/yop-center"
	YOS_SERVER_ROOT         = "https://yos.yeepay.com/yop-center"
	RSA2048                 = "RSA2048"
	YOP_PLATFORM_PUBLIC_KEY = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA6p0XWjscY+gsyqKRhw9MeLsEmhFdBRhT2emOck/F1Omw38ZWhJxh9kDfs5HzFJMrVozgU+SJFDONxs8UB0wMILKRmqfLcfClG9MyCNuJkkfm0HFQv1hRGdOvZPXj3Bckuwa7FrEXBRYUhK7vJ40afumspthmse6bs6mZxNn/mALZ2X07uznOrrc2rk41Y2HftduxZw6T4EmtWuN2x4CZ8gwSyPAW5ZzZJLQ6tZDojBK4GZTAGhnn3bg5bBsBlw2+FLkCQBuDsJVsFPiGh/b6K/+zGTvWyUcu+LUj2MejYQELDO3i2vQXVDk7lVi2/TcUYefvIcssnzsfCfjaorxsuwIDAQAB"
)

type YopRequest struct {
	// 服务地址，一般情况无需指定
	ServerRoot string
	RequestId  string
	ApiUri     string
	HttpMethod string
	AppId      string
	IsvPriKey  IsvPriKey
	// 平台公钥,一般情况无需指定
	PlatformPubKey PlatformPubKey
	// form请求的参数
	Params map[string][]string
	// json请求参数
	Content string
	// 请求头
	Headers map[string]string
	// 文件
	Files map[string]*os.File

	// 超时时间
	Timeout time.Duration
}

// NewYopRequest 创建请求
func NewYopRequest(httpMethod string, apiUri string) *YopRequest {
	return &YopRequest{HttpMethod: httpMethod, ApiUri: apiUri, Timeout: 10 * time.Second}
}

func (request *YopRequest) AddParam(name string, value any) {
	if nil == request.Params {
		request.Params = map[string][]string{}
	}
	var strValue = ToStringE(value)
	var paramArray = []string{strValue}
	request.Params[name] = paramArray
}

func (request *YopRequest) AddFile(name string, f *os.File) {
	if nil == request.Files {
		request.Files = map[string]*os.File{}
	}
	request.Files[name] = f
}

type IsvPriKey struct {
	// 密钥类型：RSA2048
	CertType string
	// 私钥值
	Value string
}

type PlatformPubKey struct {
	// 密钥类型：RSA2048
	CertType string
	// 公钥值
	Value string
}

func BuildYopRequest() *YopRequest {
	var isvPriKey = IsvPriKey{CertType: RSA2048}
	var platformCert = PlatformPubKey{Value: YOP_PLATFORM_PUBLIC_KEY, CertType: RSA2048}
	var requestId = utils.GenerateRequestID()

	return &YopRequest{RequestId: requestId, IsvPriKey: isvPriKey, PlatformPubKey: platformCert, Params: map[string][]string{}, Headers: map[string]string{}, Files: map[string]*os.File{}, Timeout: 10 * time.Second}
}

func (request *YopRequest) HandleServerRoot() {
	if len(request.ServerRoot) != 0 {
		return
	}

	if strings.HasPrefix(request.ApiUri, "/yos") {
		request.ServerRoot = YOS_SERVER_ROOT
	} else {
		request.ServerRoot = SERVER_ROOT
	}

}

var (
	errorType       = reflect.TypeOf((*error)(nil)).Elem()
	fmtStringerType = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
)

func ToStringE(i any) string {
	i = indirectToStringerOrError(i)
	switch s := i.(type) {
	case string:
		return s
	case bool:
		return strconv.FormatBool(s)
	case float64:
		return strconv.FormatFloat(s, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(s), 'f', -1, 32)
	case int:
		return strconv.Itoa(s)
	case int64:
		return strconv.FormatInt(s, 10)
	case int32:
		return strconv.Itoa(int(s))
	case int16:
		return strconv.FormatInt(int64(s), 10)
	case int8:
		return strconv.FormatInt(int64(s), 10)
	case uint:
		return strconv.FormatUint(uint64(s), 10)
	case uint64:
		return strconv.FormatUint(uint64(s), 10)
	case uint32:
		return strconv.FormatUint(uint64(s), 10)
	case uint16:
		return strconv.FormatUint(uint64(s), 10)
	case uint8:
		return strconv.FormatUint(uint64(s), 10)
	case json.Number:
		return s.String()
	case []byte:
		return string(s)
	case template.HTML:
		return string(s)
	case template.URL:
		return string(s)
	case template.JS:
		return string(s)
	case template.CSS:
		return string(s)
	case template.HTMLAttr:
		return string(s)
	case nil:
		return ""
	case fmt.Stringer:
		return s.String()
	case error:
		return s.Error()
	default:
		utils.Logger.Warnf("unable to cast %#v of type %T to string", i, i)
		return ""
	}
}

func indirectToStringerOrError(a any) any {
	if a == nil {
		return nil
	}
	v := reflect.ValueOf(a)
	for !v.Type().Implements(fmtStringerType) && !v.Type().Implements(errorType) && v.Kind() == reflect.Pointer && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

func UsePayloadForQueryParameters(yopRequest YopRequest) bool {
	var requestIsPOST = strings.Compare("POST", yopRequest.HttpMethod) == 0
	var requestHasNoPayload = len(yopRequest.Content) == 0
	return requestIsPOST && requestHasNoPayload
}
