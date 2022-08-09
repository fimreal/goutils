package net

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	firstIP, err := GetFirstIP()
	if err != nil {
		t.Fatal(err)
	}
	AllIP, err := GetAllIP()
	if err != nil {
		t.Fatal(err)
	}
	AllMacAddr, err := GetAllMacAddr()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("First IP: %q\n", firstIP)
	fmt.Printf("All IP: %q\n", AllIP)
	fmt.Printf("All MAC address: %q\n", AllMacAddr)
}
