package mbnetworkgo

import (
	"bytes"
	"encoding/gob"
	"errors"
	"strconv"
)

type Subnet struct {
	NetAddr Ip
	Suffix  uint8
}

func NewSubnet(ip Ip, suffix uint8) (Subnet, error) {
	if suffix > 32 {
		return Subnet{}, errors.New("invalid subnet suffix")
	}
	if !ip.IsValid() {
		return Subnet{}, errors.New("invalid ip")
	}
	s := Subnet{ip, suffix}
	s.NetAddr = s.CalcNetAddr()
	return s, nil
}

func (s *Subnet) CalcNetMask() Ip {
	netMask := Ip(0)
	for i := 0; i < int(s.Suffix); i++ {
		netMask |= 1 << (31 - i)
	}

	return netMask
}

func (s *Subnet) CalcNetAddr() Ip {
	s.NetAddr &= s.CalcNetMask()
	return s.NetAddr
}

func (s *Subnet) InverseNetMask() Ip {
	return ^s.CalcNetMask()
}

func (s *Subnet) Broadcast() Ip {
	return s.CalcNetAddr() | s.InverseNetMask()
}

func (s *Subnet) FirstUsable() Ip {
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
	return s.NetAddr.String() + "/" + strconv.Itoa(int(s.Suffix))
}

// TODO: add tests
func (s *Subnet) GobEncode() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(s.NetAddr); err != nil {
		return nil, err
	}
	if err := enc.Encode(s.Suffix); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
