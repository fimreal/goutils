package net

import (
	"net"

	"github.com/fimreal/goutils/ezap"
)

// 检查 IP 是否正确，是否为 IPv4 地址
// https://github.com/ip2location/ip2location-go/blob/master/iptools.go#L32
func IsIPv4(ip string) bool {
	ipaddr := net.ParseIP(ip)

	if ipaddr == nil {
		return false
	}

	v4 := ipaddr.To4()
	ezap.Info(v4)

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
	ezap.Info(v6)

	return v6 != nil
}
