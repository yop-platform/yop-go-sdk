// Package auth
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/20 2:43 PM
package auth

import (
	"crypto"
	"crypto/sha256"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/yop-platform/yop-go-sdk/yop/constants"
	"github.com/yop-platform/yop-go-sdk/yop/request"
	"github.com/yop-platform/yop-go-sdk/yop/utils"
)

var FormatISOTime = "2006-01-02T15:04:05Z"
var DEFAULT_HEADERS_TO_SIGN []string = []string{constants.YOP_APPKEY_HEADER_KEY, constants.YOP_REQUEST_ID, constants.YOP_CONTENT_SHA256}

type YopSigner interface {
	// SignRequest 请求报文签名
	SignRequest(yopRequest request.YopRequest) error

	// VerifyResponse 响应报文验签
	VerifyResponse(content string, signature string, pubKey request.PlatformPubKey) bool
}

type RsaSigner struct {
}

func (signer *RsaSigner) SignRequest(yopRequest request.YopRequest) error {
	var authString = buildAuthString(yopRequest.AppId)
	utils.Logger.Println("authString:" + authString)

	var contentHash = calculateContentHash(yopRequest)
	utils.Logger.Println("contentHash:" + contentHash)
	yopRequest.Headers[constants.YOP_CONTENT_SHA256] = contentHash

	var headerToSign = getHeaderToSign(yopRequest)
	var canonicalRequest = buildCanonicalRequest(yopRequest, authString, headerToSign)
	utils.Logger.Println("canonicalRequest:" + canonicalRequest)

	signature, err := utils.RsaSignBase64(canonicalRequest, yopRequest.IsvPriKey.Value, crypto.SHA256)
	if nil != err {
		return err
	}
	signature += "$" + "SHA256"
	utils.Logger.Println("signature:" + signature)
	var authorizationHeader = buildAuthzHeader(authString, signature, headerToSign)
	utils.Logger.Println("Authorization:" + authorizationHeader)
	yopRequest.Headers[constants.AUTHORIZATION] = authorizationHeader
	return nil
}

func (signer *RsaSigner) VerifyResponse(content string, signature string, pubKey request.PlatformPubKey) bool {
	re := regexp.MustCompile("[ \t\n]")
	content = re.ReplaceAllString(content, "")
	return utils.VerifySign(content, signature, pubKey.Value, crypto.SHA256)
}

func calculateContentHash(yopRequest request.YopRequest) string {
	var encodedParameters = ""
	if request.UsePayloadForQueryParameters(yopRequest) {
		encodedParameters = utils.GetCanonicalQueryString(yopRequest.Params)
	} else {
		encodedParameters = yopRequest.Content
	}
	utils.Logger.Println("encodedParameters:" + encodedParameters)
	return fmt.Sprintf("%x", sha256.Sum256([]byte(encodedParameters)))
}

func buildCanonicalRequest(yopRequest request.YopRequest, authString string, headerToSign []string) string {
	var canonicalQueryString = getCanonicalQueryString(yopRequest)
	var canonicalURI = getCanonicalURIPath(yopRequest.ApiUri)
	return authString + "\n" + yopRequest.HttpMethod + "\n" + canonicalURI + "\n" + canonicalQueryString + "\n" + getCanonicalHeaders(yopRequest, headerToSign)
}
func buildAuthString(appId string) string {
	var t = time.Now()
	return constants.DEFAULT_YOP_PROTOCOL_VERSION + "/" + appId + "/" + t.Format(FormatISOTime) + "/" + strconv.Itoa(constants.DEFAULT_EXPIRATION_IN_SECONDS)
}

func getCanonicalQueryString(yopRequest request.YopRequest) string {
	if request.UsePayloadForQueryParameters(yopRequest) {
		return ""
	}
	return utils.GetCanonicalQueryString(yopRequest.Params)
}

func getCanonicalURIPath(path string) string {
	if len(path) == 0 {
		return "/"
	} else if strings.HasPrefix(path, "/") {
		return utils.NormalizePath(path)
	} else {
		return "/" + utils.NormalizePath(path)
	}
}

func getHeaderToSign(yopRequest request.YopRequest) []string {
	var result []string
	for header := range DEFAULT_HEADERS_TO_SIGN {
		var value = yopRequest.Headers[DEFAULT_HEADERS_TO_SIGN[header]]
		if len(value) != 0 {
			result = append(result, DEFAULT_HEADERS_TO_SIGN[header])
		}
	}
	sort.Strings(result)
	return result
}

func getCanonicalHeaders(yopRequest request.YopRequest, headerToSign []string) string {
	var headerStrings []string
	for header := range headerToSign {
		headerStrings = append(headerStrings, utils.Normalize(headerToSign[header])+":"+utils.Normalize(yopRequest.Headers[headerToSign[header]]))
	}
	sort.Strings(headerStrings)
	return strings.Join(headerStrings, "\n")
}

func buildAuthzHeader(authString string, signature string, headerToSign []string) string {
	return "YOP-RSA2048-SHA256" + " " + authString + "/" + strings.Join(headerToSign, ";") + "/" + signature
}
