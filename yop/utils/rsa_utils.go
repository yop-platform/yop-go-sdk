package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

// RsaSignBase64 base64UrlEncode签名
func RsaSignBase64(content string, privateKey string, hash crypto.Hash) (string, error) {
	signature, err := Sign(content, privateKey, hash)
	if err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(signature), nil
}

// VerifySign 验签
func VerifySign(content string, signature string, pubKey string, hash crypto.Hash) bool {
	hashed := sha256.Sum256([]byte(content))
	pubKey = FormatPemKey(pubKey, "PUBLIC KEY")
	publicKey, err := ParsePublicKey(pubKey)
	if err != nil {
		return false
	}
	sig, _ := base64.RawURLEncoding.DecodeString(signature)
	err = rsa.VerifyPKCS1v15(publicKey, hash, hashed[:], sig)
	if err != nil {
		return false
	}
	return true
}

// Sign rsa签名
func Sign(content string, privateKey string, hash crypto.Hash) ([]byte, error) {
	shaNew := sha256.New()
	shaNew.Write([]byte(content))
	hashed := shaNew.Sum(nil)

	priKey, err := ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey.(*rsa.PrivateKey), hash, hashed)

	return signature, err
}

func ParsePrivateKey(privateKey string) (any, error) {
	privateKey = FormatPemKey(privateKey, "PRIVATE KEY")
	// 2、解码私钥字节，生成加密对象
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}

	// 3、解析DER编码的私钥，生成私钥对象
	priKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return priKey, nil
}

func ParsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*rsa.PublicKey), nil
}

// FormatPemKey /**
func FormatPemKey(yopFormKey string, pemHeader string) string {
	var sb = strings.Builder{}
	sb.WriteString("-----BEGIN ")
	sb.WriteString(pemHeader)
	sb.WriteString("-----\n")
	for i := 0; i < len(yopFormKey); i++ {
		sb.WriteString(string([]rune(yopFormKey)[i]))
		if (i+1)%64 == 0 {
			sb.WriteString("\n")
		}
	}
	sb.WriteString("\n-----END ")
	sb.WriteString(pemHeader)
	sb.WriteString("-----\n")
	return sb.String()
}
