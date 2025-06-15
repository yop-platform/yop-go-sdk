// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/29 10:09 AM
package utils

import "testing"

func TestParseToJson(t *testing.T) {
	var params = map[string]any{}
	params["p1"] = "11"
	params["p2"] = 2
	str := ParseToJsonStr(params)
	if len(str) == 0 {
		t.Fatal("parse failed")
	}
	t.Log(str)
}
