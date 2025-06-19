// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/20 4:15 PM
package utils

import (
	"net/url"
	"sort"
	"strings"
)

func NormalizePath(path string) string {

	return strings.ReplaceAll(Normalize(path), "%2F", "/")
}

func Normalize(value string) string {
	var firstEncodeStr = url.QueryEscape(value)
	return encodeSpecialChar(firstEncodeStr)
}

func encodeSpecialChar(str string) string {
	// 空格
	str = strings.ReplaceAll(str, "+", "%20")
	return str
}

func EncodeParameters(params map[string][]string, forSign bool) string {
	if len(params) == 0 {
		return ""
	}
	var encodedNameValuePair []string
	for k, v := range params {
		for i := range v {
			encodedNameValuePair = append(encodedNameValuePair, toNameValuePair(k, v[i], forSign))
		}
	}
	return strings.Join(encodedNameValuePair, "&")
}

func toNameValuePair(paramName string, paramValue string, forSign bool) string {
	val := paramValue
	if !forSign {
		val = Normalize(paramValue)
	}
	return Normalize(paramName) + "=" + Normalize(val)
}

func GetCanonicalQueryString(params map[string][]string) string {
	if len(params) == 0 {
		return ""
	}

	var parameterStrings []string

	for k, v := range params {
		if len(v) == 0 {
			parameterStrings = append(parameterStrings, Normalize(k)+"=")
		} else {
			for i := range v {
				parameterStrings = append(parameterStrings, Normalize(k)+"="+Normalize(v[i]))
			}
		}

	}
	sort.Strings(parameterStrings)
	return strings.Join(parameterStrings, "&")
}
