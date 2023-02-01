package mbnetworkgo

import (
	"errors"
	"strconv"
)

type subnet struct {
	netAddr Ip
	suffix  uint8
}

func NewSubnet(ip Ip, suffix uint8) (subnet, error) {
	if suffix > 32 {
		return subnet{}, errors.New("invalid subnet suffix")
	}
	if !ip.IsValid() {
		return subnet{}, errors.New("invalid ip")
	}
	s := subnet{ip, suffix}
	s.netAddr = s.CalcNetAddr()
	return s, nil
}

func (s *subnet) CalcNetMask() Ip {
	netMask := Ip(0)
	for i := 0; i < int(s.suffix); i++ {
		netMask |= 1 << (31 - i)
	}

	return netMask
}

func (s *subnet) CalcNetAddr() Ip {
	s.netAddr &= s.CalcNetMask()
	return s.netAddr
}

func (s *subnet) InverseNetMask() Ip {
	return ^s.CalcNetMask()
}

func (s *subnet) Broadcast() Ip {
	return s.CalcNetAddr() | s.InverseNetMask()
}

func (s *subnet)	FirstUsable() Ip {
	return s.CalcNetAddr() + 1
}

func (s *subnet) LastUsable() Ip {
	return s.Broadcast() - 1
}

func (s *subnet) PossibleHosts() uint32 {
	return uint32(s.LastUsable() - s.FirstUsable() + 1)
}

func (s *subnet) Contains(ip Ip) bool {
	return ip >= s.FirstUsable() && ip <= s.LastUsable()
}

func (s subnet) String() string {
	return s.netAddr.String() + "/" + strconv.Itoa(int(s.suffix))
}