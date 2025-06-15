package utils

import (
	"fmt"
	"github.com/gofrs/uuid/v5"
	"os"
	"time"
)

// GenerateRequestID 生成请求ID，如果UUID生成失败则使用备用方案
func GenerateRequestID() string {
	uuidV4, err := uuid.NewV4()
	if err != nil {
		Logger.Println("Failed to generate UUID, using fallback: " + err.Error())
		return fmt.Sprintf("fallback-%d-%d", time.Now().UnixNano(), os.Getpid())
	}
	return uuidV4.String()
}
