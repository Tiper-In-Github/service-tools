package utils_test

import (
	"testing"

	"service-tools/utils" // 替换为你的模块路径
)

// TestGetLocalIP_Success 测试正常情况：能否成功获取本地非回环IPv4地址
func TestGetLocalIP_Success(t *testing.T) {
	ip, err := utils.GetLocalIP()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if ip == nil {
		t.Fatal("expected non-nil IP address")
	}

	if ip.To4() == nil {
		t.Errorf("expected IPv4 address, got IPv6 or invalid: %v", ip)
	}

	if ip.IsLoopback() {
		t.Errorf("expected non-loopback address, got loopback: %v", ip)
	}

	t.Logf("Detected local IP: %v", ip)
}
