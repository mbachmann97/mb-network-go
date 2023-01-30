package main

import (
	"fmt"
	"os"
)

func main() {
	ip, err := NewIpFromString("192.168.1.189")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(ip)

	subnet, err := NewSubnet(ip, 22)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	fmt.Println(subnet.CalcNetMask())
	fmt.Println(subnet.CalcNetAddr())
	fmt.Println(subnet.PossibleHostsCount())
	fmt.Println(subnet.InverseNetMask())
	fmt.Println(subnet.Broadcast())
	fmt.Println(subnet)
}
