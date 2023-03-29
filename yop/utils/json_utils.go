// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/29 9:56 AM
package utils

import "encoding/json"

func ParseToJsonStr(params map[string]any) string {
	marshal, _ := json.Marshal(params)
	return string(marshal)
}
