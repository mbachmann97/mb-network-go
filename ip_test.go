package main

import "testing"

func TestNewIpFromStringWithWrongOctetCountWillFail(t *testing.T) {
	_, err := newIpFromString("10.10.0.10.10")
	if err == nil {
		t.Error("Expected octet count error got", err)
	}
}

func TestNewIpFromStringWithInvalidOctetsWillFail(t *testing.T) {
	_, err := newIpFromString("258.10.0.10")
	if err == nil {
		t.Error("Expected octet(s) invalid error got", err)
	}
}

func TestNewIpFromStringEmptyStringWillFail(t *testing.T) {
	_, err := newIpFromString("")
	if err == nil {
		t.Error("Expected octet count error got", err)
	}
}

func TestIpToString(t *testing.T) {
	ip, _ := newIpFromString("192.168.1.1")
	if ip.String() != "192.168.1.1" {
		t.Error("Ip to string conversion failed")
	}
}