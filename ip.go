package main

import (
	"errors"
	"strconv"
	"strings"
)

type ip uint32

func newIpFromString(s string) (ip, error) {
	var packaged ip
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
		packaged = ip(octets[0]<<24 | octets[1]<<16 | octets[2]<<8 | octets[3])

		return packaged, nil
	} else {
		return 0, errors.New("octet count not equal to 4")
	}
}

func (i ip) String() string {
	return strconv.Itoa(int(i>>24)) + "." + strconv.Itoa(int(i>>16&0xFF)) + "." + strconv.Itoa(int(i>>8&0xFF)) + "." + strconv.Itoa(int(i&0xFF))
}