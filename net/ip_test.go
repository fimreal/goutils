package net

import "testing"

func TestIsIPv4(t *testing.T) {
	ip := "1.1.1.1"
	if !IsIPv4(ip) {
		t.Error("err")
	}

	ip = "epurs.com"
	if IsIPv4(ip) {
		t.Error("err")
	}
	ip = "1.225.250.1"
	if IsIPv4(ip) {
		t.Error("err")
	}
	ip = "01.225.223.1"
	if IsIPv4(ip) {
		t.Error("err")
	}
}

func TestIsIPv6(t *testing.T) {
	ip := "fe80::742d:e8ff:feb6:cc08"
	if !IsIPv6(ip) {
		t.Error("err")
	}

	ip = "epurs6.com"
	if IsIPv4(ip) {
		t.Error("err")
	}
}
