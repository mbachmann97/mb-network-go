package mbnetworkgo

import (
	"errors"
	"strconv"
)

type Subnet struct {
	netAddr Ip
	suffix  uint8
}

func NewSubnet(ip Ip, suffix uint8) (Subnet, error) {
	if suffix > 32 {
		return Subnet{}, errors.New("invalid subnet suffix")
	}
	if !ip.IsValid() {
		return Subnet{}, errors.New("invalid ip")
	}
	s := Subnet{ip, suffix}
	s.netAddr = s.CalcNetAddr()
	return s, nil
}

func (s *Subnet) CalcNetMask() Ip {
	netMask := Ip(0)
	for i := 0; i < int(s.suffix); i++ {
		netMask |= 1 << (31 - i)
	}

	return netMask
}

func (s *Subnet) CalcNetAddr() Ip {
	s.netAddr &= s.CalcNetMask()
	return s.netAddr
}

func (s *Subnet) InverseNetMask() Ip {
	return ^s.CalcNetMask()
}

func (s *Subnet) Broadcast() Ip {
	return s.CalcNetAddr() | s.InverseNetMask()
}

func (s *Subnet)	FirstUsable() Ip {
	return s.CalcNetAddr() + 1
}

func (s *Subnet) LastUsable() Ip {
	return s.Broadcast() - 1
}

func (s *Subnet) PossibleHosts() uint32 {
	return uint32(s.LastUsable() - s.FirstUsable() + 1)
}

func (s *Subnet) Contains(ip Ip) bool {
	return ip >= s.FirstUsable() && ip <= s.LastUsable()
}

func (s Subnet) String() string {
	return s.netAddr.String() + "/" + strconv.Itoa(int(s.suffix))
}