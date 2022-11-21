package net

import (
	"net"
	"strings"
)

// 检查 IP 是否正确，是否为 IPv4 地址
// https://github.com/ip2location/ip2location-go/blob/master/iptools.go#L32
func IsIPv4(ip string) bool {
	ipaddr := net.ParseIP(ip)

	if ipaddr == nil {
		return false
	}

	v4 := ipaddr.To4()
	// ezap.Info(v4)

	return v4 != nil
}

// 检查 IP 是否正确，是否为 IPv6 地址
// https://github.com/ip2location/ip2location-go/blob/master/iptools.go#L49
func IsIPv6(ip string) bool {
	if IsIPv4(ip) {
		return false
	}

	ipaddr := net.ParseIP(ip)

	if ipaddr == nil {
		return false
	}

	v6 := ipaddr.To16()
	// ezap.Info(v6)

	return v6 != nil
}

func IsLanIPv4(ip string) bool {
	if !IsIPv4(ip) {
		return false
	}

	lanex := []string{"10.", "192.168.", "172.16.", "172.17.", "172.18.", "172.19.", "172.20.", "172.21.", "172.22.", "172.23.", "172.24.", "172.25.", "172.26.", "172.27.", "172.28.", "172.29.", "172.30.", "172.31."}

	for _, lan := range lanex {
		if strings.HasPrefix(ip, lan) {
			return true
		}
	}
	return false
}
