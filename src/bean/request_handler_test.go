package main

import (
	"testing"
)

func Test_ListAuthenticator(t *testing.T) {
	t.Parallel()
	handler := NewRequestHandler("../../configs")
	data, err := handler.ListAuthenticator([]string{"www.baidu.com"})
	if err != nil {
		t.Error("ListAuthenticator Error")
	}
	if len(data) != 2 || data[0] != "192.168.0.4" && data[1] != "127.0.0.1" {
		t.Error("ListAuthenticator data Error")
	}
}

func Test_IP(t *testing.T) {
	t.Parallel()
	handler := NewRequestHandler("../../configs")
	yes_3 := false
	yes_4 := false
	for i := 0; i < 100; i++ {
		data, err := handler.IP([]string{"www.baidu.com"})
		if err != nil {
			t.Error("Test_IP Error")
		} else if data[0] == "192.168.0.3" {
			yes_3 = true
		} else if data[0] == "192.168.0.4" {
			yes_4 = true
		} else {
			t.Error("Unexpected Result while Test_IP")
		}
	}
	if !yes_3 || !yes_4 {
		t.Error("IP Random error")
	}
}

func Test_ScoreIP(t *testing.T) {
	t.Parallel()
	handler := NewRequestHandler("../../configs")
	data, err := handler.ScoreIP([]string{"www.baidu.com", "192.168.0.4", "-10"})
	if err != nil {
		t.Error("Test_ScoreIP Error")
	}
	if len(data) != 2 || data[0] != "192.168.0.4" && data[1] != "90" {
		t.Error("Test_ScoreIP data Error")
	}

	for i := 0; i < 10; i++ {
		data, err = handler.IP([]string{"www.baidu.com"})
		if len(data) != 1 && data[0] != "192.168.0.3" {
			t.Error("Test_ScoreIP data Error")
		}
	}
}

func Test_Timeout(t *testing.T) {
	t.Parallel()
	handler := NewRequestHandler("../../configs")
	data, err := handler.Timeout([]string{"www.baidu.com"})
	if err != nil {
		t.Error("Test_Timeout Error")
	}
	if len(data) != 3 || data[0] != "3500" && data[1] != "150" && data[2] != "200" {
		t.Error("Test_Timeout data Error:", data)
	}
}
