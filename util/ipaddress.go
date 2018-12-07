package util

import (
	"bytes"
	"net"
	"net/http"
	"strings"
)

type ipRange struct {
	start net.IP
	end   net.IP
}

func newIPRange(start, end string) *ipRange {
	return &ipRange{
		start: net.ParseIP(start),
		end:   net.ParseIP(end),
	}
}

func (r *ipRange) contains(ipAddress net.IP) bool {
	return bytes.Compare(ipAddress, r.start) >= 0 && bytes.Compare(ipAddress, r.end) < 0
}

var privateRanges = []*ipRange{
	newIPRange("127.0.0.0", "127.255.255.255"),
	newIPRange("10.0.0.0", "10.255.255.255"),
	newIPRange("100.64.0.0", "100.127.255.255"),
	newIPRange("172.16.0.0", "172.31.255.255"),
	newIPRange("192.0.0.0", "192.0.0.255"),
	newIPRange("192.168.0.0", "192.168.255.255"),
	newIPRange("198.18.0.0", "198.19.255.255"),
}

// isPrivateSubnet checks to see if this ip is in a private subnet
func isPrivateSubnet(ipAddress net.IP) bool {
	if ipCheck := ipAddress.To4(); ipCheck != nil {
		for _, r := range privateRanges {
			if r.contains(ipAddress) {
				return true
			}
		}
	}
	return false
}

// GetRemoteIP attempts to determine the "real" source of the request by
// examining the X-Forwarded-For and X-Real-Ip headers. If an address cannot be
// found in the headers, the request's RemoteAddr is returned.
func GetRemoteIP(r *http.Request) string {
	for _, h := range []string{"X-Forwarded-For", "X-Real-Ip"} {
		addresses := strings.Split(r.Header.Get(h), ",")
		// march from right to left until we get a public address
		// that will be the address right before our proxy.
		for i := len(addresses) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(addresses[i])
			// header can contain spaces too, strip those out.
			realIP := net.ParseIP(ip)
			if realIP.IsGlobalUnicast() && !isPrivateSubnet(realIP) {
				return ip
			}
		}
	}
	return r.RemoteAddr
}
