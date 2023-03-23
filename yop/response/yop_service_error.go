// Package response
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/22 11:09 PM
package response

type YopServiceError struct {
	RequestId  string
	Code       string
	Message    string
	SubCode    string
	SubMessage string
	DocUrl     string
}

func (err *YopServiceError) Error() string {
	return err.Message + "(Error Code: " + err.Code + "; Sub Code: " + err.SubCode + "; Sub Message: " + err.SubMessage + "; Request ID: " + err.RequestId + "; docUrl: " + err.DocUrl + ")"
}
