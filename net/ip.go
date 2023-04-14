package net

import (
	"net"
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

// 判断是否为保留网络地址
// https://en.wikipedia.org/wiki/Reserved_IP_addresses
func IsLanIPv4(ip string) bool {
	if !IsIPv4(ip) {
		return false
	}
	netIP := net.ParseIP(ip)
	revNet := []string{
		"0.0.0.0/8",
		"10.0.0.0/8",
		"100.64.0.0/10",
		"127.0.0.0/8",
		"169.254.0.0/16",
		"172.16.0.0/12",
		"192.0.0.0/24",
		"192.0.2.0/24",
		"192.88.99.0/24",
		"192.168.0.0/16",
		"198.18.0.0/15",
		"198.51.100.0/24",
		"203.0.113.0/24",
		"224.0.0.0/4",
		"240.0.0.0/4",
		"255.255.255.255/32",
	}

	for _, cidr := range revNet {
		_, ipNet, _ := net.ParseCIDR(cidr)
		if ipNet.Contains(netIP) {
			return true
		}
	}
	return false
}
