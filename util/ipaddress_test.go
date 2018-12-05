package util

import (
	"fmt"
	"net/http"
	"testing"
)

func newRequestWithHeader(header, value string) *http.Request {
	return &http.Request{
		Header: map[string][]string{
			header: []string{value},
		},
	}
}

func TestGetRemoteIP(t *testing.T) {
	rem := "62.62.62.62"
	cases := []struct {
		req *http.Request
		ip  string
	}{
		{newRequestWithHeader("X-Forwarded-For", rem), rem},
		{newRequestWithHeader("X-Real-Ip", rem), rem},
		{newRequestWithHeader("X-Forwarded-For", fmt.Sprintf("10.10.0.5, 98.132.64.8, %s, 192.168.32.8", rem)), rem},
		{&http.Request{RemoteAddr: rem}, rem},
	}
	for _, c := range cases {
		if ip := GetRemoteIP(c.req); ip != c.ip {
			t.Errorf("expected %s but got %s", c.ip, ip)
		}
	}
}
