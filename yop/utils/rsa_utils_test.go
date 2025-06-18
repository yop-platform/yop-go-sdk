// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/14 11:23 AM
package utils

import (
	"crypto"
	"testing"
)

func TestRsaSignBase64(t *testing.T) {
	var priKey = "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDBHBdHbQXsPT+EpAhLA9k2Q5O8GLCAUFLWYB57Uhc4ZNa2YUhjrTFvFZMFQuMjaVgdmFGTvqfGYUQBRldHFhf9kuXf5LPb+m0BJ/R5AWCyTcX7DHouoGODMfxkCZrimILwYWDkwhYTHr5hEV58nGRQtHOIVB5a4i/y4Z1vvX1MIjA+8OJ3zpaxXKkj+46OtfmjloUPGFSzz+rqrRRtMqYePLkWZ0J+CmIXM1Kwl/kgYUq/YGSYy9Q5vTojN9WBKzk7euOoCcsWtRrBQysdyM3yDPXjhnXx7G0nh07hSUh+rMDZ7Zst0lVQol/7kPXzNwh6eUmRGY9lfruIMS5kg1ydAgMBAAECggEAD4yQf0rTCEOiQq7mkAu+SLVGRwYB6EMPeH2C1tE0V3EfLM5GgugmK9ij3u+U1HweATwLjYbzgXDBhgzA6FNqGRvj8JQ8u0C92DL8Z2XqAFFs2JsXl3uIp761oOR5GTfIi0x7/c928ZEvKSe54PTCyxDMoLSNQSonTDpIb//k/+U4xEOQ1mjlSvlOM5ic7/kdw+G+aP/Hk/T6kg/vIblWQHx8SB3WYpLb/R6oPO+05X+zcQ+vVX1TrQ/amDp6/PouWjTF5hf48JEBdM8+xJzUwnalrG9U7pChfyGAOXQT1fbDdywBJXt6pZsT/mz1RkUC5Uto5/aVQGIDD+IPm/ZDbQKBgQDEwF09sjUb5hHEdmG28RmMf4E0JCOEzCvxiUpovobymqapLM5bf2oLNXqGenEAMbfFQatJFVKx6YBZwFIj/xzQJt8fL/jRzlbLijaANP+1JacvTsfXKBXS888FN3rkKisTlhmYXI+4EwA1wbcRkLDH3vezdVCi9cszQ9HvwkVfOwKBgQD7QvzD3pirXIJ64JizWTS4MJMko3CWepsq9UZ5uyoHWh7tSz86H/2y0FK10YpJEJeGtyPXlnU+uQwjYMJRPLlNv8180pjCJX2ZTW2drB2vOJvormhMhDIYAZtPAHu2dajzdy4VRuvFTtH4FpW/KjAJrTLK3ze3K95ACYVBJ8EmBwKBgB0K8DiNN724hmLjvqTMjiLpJ19U/lE5+jqbM3qmtTDWl0ddr9BdzH9/E2kKZefLbv8VJH2TQjO07hdRhk599/jZ5BGseSQvOyysaEMgj6ZjunwHOwSNjDspdiOk/uTzPIyVmY2eDDD1zRAiWi2jmBTI2vOIm7CSa75TgofLu4XFAoGBAJrFM4+vYNlFXbY0/LqU+21ttmV+K471rPj0Jto7GPN4Zs6CaEr0g8COpDQNA6JoDv5Td0eIDWZ6c+ii5G9H+VjUCc6WprQIhepVkGzsJUjWlOrp66MeVwEElFdAk/PbXBvEUOWYTwi1uY6Y0trzMK31OvFOODKjWf6WHrf4tfgnAoGAOri6bX2D/zqpJT3mJ5MIVJJbn4D4Idx+TCUaVRSY1rBp+Y2ofW1W8ktu7xPO9/LwVQR7kJeosEBAFGTmGqll033ywu5+8X8J1bw6HCghkI0yHW752sOdfl30kXi3Ds8tQsvSEHRfnPb8yvWve2srZb9ubwOvpI0PtOIujZP4fYI="
	var content = "{\"result\":{\"requestId\":\"requestId\",\"errorMsg\":\"exception.record.not.found.transferDomesticOrder|merchantId:[null],requestId:[requestId]\",\"status\":\"FAILED\"}}"
	signature, error := RsaSignBase64(content, priKey, crypto.SHA256)
	if nil != error {
		t.Fatal("exception happened")
	}

	if len(signature) == 0 {
		t.Fatal("signature is empty")
	}
	t.Log(signature)
}

func TestVerifySign(t *testing.T) {
	var pubKey = "MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwRwXR20F7D0/hKQISwPZNkOTvBiwgFBS1mAee1IXOGTWtmFIY60xbxWTBULjI2lYHZhRk76nxmFEAUZXRxYX/ZLl3+Sz2/ptASf0eQFgsk3F+wx6LqBjgzH8ZAma4piC8GFg5MIWEx6+YRFefJxkULRziFQeWuIv8uGdb719TCIwPvDid86WsVypI/uOjrX5o5aFDxhUs8/q6q0UbTKmHjy5FmdCfgpiFzNSsJf5IGFKv2BkmMvUOb06IzfVgSs5O3rjqAnLFrUawUMrHcjN8gz144Z18extJ4dO4UlIfqzA2e2bLdJVUKJf+5D18zcIenlJkRmPZX67iDEuZINcnQIDAQAB"
	var signature = "glTZg6lLl6oV4Ho15fAUegcVILlTwYJkbZO_Iz8AYUKTZ_1JP4AqAqSdm3GqjaukoNrDkxPGv2WW8plxYxtzsXjkzWiCMth5aShHgA7a9SXW0jfo365KPyVj0zFO2QIV9odHEnY1apwcAxvr54j4d5SHoC3vKUczZ20txTsNjcG9ifi1AoJhblILxKL2NO0tdIzTMQCRaBdOXUOdnL7RgP1qPew5yJT4e1QdtTjkirCKJurm4SumOA3Uroz-G-9MUZgiTkU4RXrEvu-rJPlqfJPsITYoWLsuPy1Gfne_5j-IgChXpoHacI0s-NlzKmyjsFt3-5aUYDd0cFw58ErUXw"
	var data = "{\"result\":{\"requestId\":\"requestId\",\"errorMsg\":\"exception.record.not.found.transferDomesticOrder|merchantId:[null],requestId:[requestId]\",\"status\":\"FAILED\"}}"
	if !VerifySign(data, signature, pubKey, crypto.SHA256) {
		t.Fatal("verify failed")
	}
}
