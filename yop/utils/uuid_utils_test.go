package utils

import (
	"strings"
	"testing"
)

func TestGenerateRequestID(t *testing.T) {
	// 测试UUID生成功能
	requestID := GenerateRequestID()

	// 验证生成的ID不为空
	if requestID == "" {
		t.Fatal("Generated request ID should not be empty")
	}

	// 验证ID格式（UUID或备用格式）
	if strings.Contains(requestID, "fallback-") {
		// 如果是备用格式，应该包含时间戳和进程ID
		parts := strings.Split(requestID, "-")
		if len(parts) < 3 {
			t.Errorf("Fallback request ID format is incorrect: %s", requestID)
		}
		t.Logf("Using fallback request ID: %s", requestID)
	} else {
		// 如果是UUID格式，应该包含连字符
		if !strings.Contains(requestID, "-") {
			t.Errorf("UUID format is incorrect: %s", requestID)
		}
		t.Logf("Generated UUID: %s", requestID)
	}

	// 测试多次生成的ID应该不同
	requestID2 := GenerateRequestID()
	if requestID == requestID2 {
		t.Errorf("Multiple calls should generate different IDs: %s == %s", requestID, requestID2)
	}
}

func TestGenerateRequestIDUniqueness(t *testing.T) {
	// 测试ID唯一性
	ids := make(map[string]bool)
	for i := 0; i < 100; i++ {
		id := GenerateRequestID()
		if ids[id] {
			t.Errorf("Duplicate ID generated: %s", id)
		}
		ids[id] = true
	}
	t.Logf("Generated %d unique IDs", len(ids))
}
