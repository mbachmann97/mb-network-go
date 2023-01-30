package main

import "testing"

func TestNewIpFromString(t *testing.T) {
	t.Run("NewIpFromString with wrong octet count will fail", func(t *testing.T) {
		_, err := NewIpFromString("10.10.0.10.10")
	if err == nil {
		t.Error("Expected octet count error got", err)
	}
	})

	t.Run("NewIpFromString with invalid octets will fail", func(t *testing.T) {
		_, err := NewIpFromString("258.10.0.10")
	if err == nil {
		t.Error("Expected octet(s) invalid error got", err)
	}
	})

	t.Run("NewIpFromString empty string will fail", func(t *testing.T) {
		_, err := NewIpFromString("")
	if err == nil {
		t.Error("Expected octet count error got", err)
	}
	})
}

func TestIpToString(t *testing.T) {
	ip, _ := NewIpFromString("192.168.1.1")
	if ip.String() != "192.168.1.1" {
		t.Error("Ip to string conversion failed")
	}
}