// Package utils
// Copyright: Copyright (c) 2020<br>
// Company: 易宝支付(YeePay)<br>
// @author    : yunmei.wu
// @time      : 2023/3/15 2:22 PM
package utils

import (
	"crypto"
	"crypto/aes"
	"encoding/base64"
	"errors"
	"strings"
)

// DecryptCallback 解密回调通知内容
func DecryptCallback(platformPubKey string, isvPriKey string, callBack string) (string, error) {
	cipherText := strings.Split(callBack, "$")
	if len(cipherText) != 4 {
		return "", errors.New("response invalid")
	}
	randomKey, err := RsaDecrypt(isvPriKey, cipherText[0])
	if err != nil {
		Logger.Println("random key rsa error ", err)
		return "", err
	}
	cipherBytes := base64Decode(cipherText[1])
	body := string(AesDecryptECB(cipherBytes, randomKey))
	dollarPosition := strings.LastIndex(body, "$")
	signature := strings.TrimSpace(body[dollarPosition+1:])
	body = body[:dollarPosition]

	if !VerifySign(body, signature, platformPubKey, crypto.SHA256) {
		return "", errors.New("rsa sign verify fail")
	}
	return body, nil
}

// 解析非标 base64_encode
func base64Decode(b string) []byte {
	b = strings.Replace(b, "-", "+", -1)
	b = strings.Replace(b, "_", "/", -1)
	r, err := base64.RawStdEncoding.DecodeString(b)
	if err != nil {
		Logger.Println("base64 decode error ", err)
		return nil
	}

	return r
}
func AesDecryptECB(encrypted []byte, key []byte) (decrypted []byte) {
	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
