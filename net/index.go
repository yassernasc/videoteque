package net

import "net"

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
