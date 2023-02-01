package mbnetworkgo

import (
	"errors"
	"strconv"
	"strings"
)

type Ip uint32

func NewIpFromString(s string) (Ip, error) {
	var packaged Ip
	stringOctets := strings.Split(s, ".")
	if len(stringOctets) == 4 {
		octets := make([]uint32, 4)
		for i, octet := range stringOctets {
			o, err := strconv.ParseUint(octet, 10, 8)
			if err != nil || o > 255 {
				return 0, errors.New("one or more octets are not valid")
			}
			octets[i] = uint32(o)
		}
		packaged = Ip(octets[0]<<24 | octets[1]<<16 | octets[2]<<8 | octets[3])

		return packaged, nil
	} else {
		return 0, errors.New("octet count not equal to 4")
	}
}

func (i Ip) IsValid() bool {
	octets := []uint32{uint32(i >> 24), uint32(i >> 16 & 0xFF), uint32(i >> 8 & 0xFF), uint32(i & 0xFF)}
	for _, octet := range octets {
		if octet > 255 {
			return false
		}
	}
	return true
}

func (i Ip) String() string {
	return strconv.Itoa(int(i>>24)) + "." + strconv.Itoa(int(i>>16&0xFF)) + "." + strconv.Itoa(int(i>>8&0xFF)) + "." + strconv.Itoa(int(i&0xFF))
}