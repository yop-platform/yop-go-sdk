package request

import (
	"testing"
	"time"
)

func TestNewYopRequestTimeout(t *testing.T) {
	// 测试NewYopRequest的默认超时时间
	req := NewYopRequest("GET", "/test")

	expectedTimeout := 60 * time.Second
	if req.Timeout != expectedTimeout {
		t.Errorf("Expected timeout %v, got %v", expectedTimeout, req.Timeout)
	}

	t.Logf("Default timeout correctly set to %v", req.Timeout)
}

func TestBuildYopRequestTimeout(t *testing.T) {
	// 测试BuildYopRequest的默认超时时间
	req := BuildYopRequest()

	expectedTimeout := 60 * time.Second
	if req.Timeout != expectedTimeout {
		t.Errorf("Expected timeout %v, got %v", expectedTimeout, req.Timeout)
	}

	t.Logf("BuildYopRequest timeout correctly set to %v", req.Timeout)
}

func TestYopRequestCustomTimeout(t *testing.T) {
	// 测试自定义超时时间
	req := NewYopRequest("POST", "/test")
	customTimeout := 30 * time.Second
	req.Timeout = customTimeout

	if req.Timeout != customTimeout {
		t.Errorf("Expected custom timeout %v, got %v", customTimeout, req.Timeout)
	}

	t.Logf("Custom timeout correctly set to %v", req.Timeout)
}

func TestUsePayloadForQueryParameters(t *testing.T) {
	// 测试UsePayloadForQueryParameters函数
	tests := []struct {
		name     string
		method   string
		content  string
		expected bool
	}{
		{"POST with no content", "POST", "", true},
		{"POST with content", "POST", "some content", false},
		{"GET with no content", "GET", "", false},
		{"GET with content", "GET", "some content", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := YopRequest{
				HttpMethod: tt.method,
				Content:    tt.content,
			}

			result := UsePayloadForQueryParameters(req)
			if result != tt.expected {
				t.Errorf("UsePayloadForQueryParameters() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
