package mbnetworkgo

import "testing"

func TestNewSubnet(t *testing.T) {
	t.Run("NewSubnet with invalid suffix will fail", func(t *testing.T) {
		ip, _ := NewIpFromString("192.160.1.1")
		_, err := NewSubnet(ip, 33)
		if err == nil {
			t.Error("Expected suffix invalid error got", err)
		}
	})
	t.Run("NewSubnet sets netAddr correctly", func(t *testing.T) {
		ip, _ := NewIpFromString("192.160.1.23")
		s, _ := NewSubnet(ip, 24)
		if s.netAddr != 0xc0a00100 {
			t.Error("Expected suffix 192.160.1.0 got", s.netAddr)
		}
	})
}

func TestCalcNetMask(t *testing.T) {
	t.Run("CalcNetMask with suffix 24 will return 255.255.255.0", func(t *testing.T) {
		s := subnet{netAddr: 0, suffix: 24}
		if s.CalcNetMask() != 0xffffff00 {
			t.Error("Expected the netmask to be 255.255.255.0 got", s.CalcNetMask())
		}
	})
}

func TestCalcNetAddr(t *testing.T) {
	t.Run("CalcNetAddr with ip 192.168.23.56 and suffix 16 will return 192.168.0.0", func(t *testing.T) {
		ip, _ := NewIpFromString("192.168.23.56")
		s, _ := NewSubnet(ip, 16)
		if s.CalcNetAddr() != 0xc0a80000 {
			t.Error("Expected the netaddr to be 192.168.0.0 got", s.CalcNetAddr())
		}
	})
}

func TestInverseNetMask(t *testing.T) {
	t.Run("255.252.0.0 is 0.2.255.255 inversed", func(t *testing.T) {
		s := subnet{netAddr: 0, suffix: 14}
		if s.InverseNetMask() != 0x0003ffff {
			t.Error("Expected the inverse netmask to be 0.2.255.255 got", s.InverseNetMask())
		}
	})
}

func TestBroadcast(t *testing.T) {
	t.Run("Broadcast of 192.168.0.0/16 is 192.168.255.255", func(t *testing.T) {
		ip, _ := NewIpFromString("192.168.0.0")
		s, _ := NewSubnet(ip, 16)
		if s.Broadcast() != 0xc0a8ffff {
			t.Error("Expected the broadcast to be 192.168.255.255 got", s.Broadcast())
		}
	})
}

func TestPossibleHostsCount(t *testing.T) {
	t.Run("PossibleHostsCount of 10.22.1.32/8 is 16777214", func(t *testing.T) {
		ip, _ := NewIpFromString("10.22.1.32")
		s, _ := NewSubnet(ip, 8)
		if s.PossibleHosts() != 16777214 {
			t.Error("Expected the possible hosts count to be 16777214 got", s.PossibleHosts())
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("192.168.1.77/26 contains 192.168.1.65", func(t *testing.T) {
		ip, _ := NewIpFromString("192.168.1.77")
		s, _ := NewSubnet(ip, 26)
		target, _ := NewIpFromString("192.168.1.65")
		if !s.Contains(target) {
			t.Error("Expected the subnet to contain 192.168.1.65")
		}
	})

	t.Run("10.1.2.4/8 does not contain 192.160.1.23", func(t *testing.T) {
		ip, _ := NewIpFromString("10.1.2.4")
		s, _ := NewSubnet(ip, 8)
		target, _ := NewIpFromString("192.160.1.23")
		if s.Contains(target) {
			t.Error("Expected the subnet to not contain 192.160.1.23")
		}
	})
}

func TestString(t *testing.T) {
	t.Run("String of ip 8.8.8.8 and suffix 8 is 8.8.8.8/8", func(t *testing.T) {
		s := subnet{netAddr: 0x08080808, suffix: 8}
		if s.String() != "8.8.8.8/8" {
			t.Error("Expected the subnet to be 8.8.8.8/8 got", s.String())
		}
	})
}