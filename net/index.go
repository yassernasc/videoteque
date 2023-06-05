package net

import (
	"net"
	"strings"
)

func LocalIp() string {
	addrs, _ := net.InterfaceAddrs()

	for _, addr := range addrs {
		ip, _, _ := net.ParseCIDR(addr.String())
		if i := ip.To4(); i != nil && !ip.IsLoopback() {
			return i.String()
		}
	}

	return ""
}

func AvoidLocalhostNotation(url string) string {
	notations := [...]string{"localhost", "127.0.0.1"}
	for _, n := range notations {
		if strings.Contains(url, n) {
			return strings.Replace(url, n, LocalIp(), -1)
		}
	}
	return url
}
