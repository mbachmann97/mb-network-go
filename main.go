package main

import (
	"fmt"
	"os"
)

func main() {
	ip, err := newIpFromString("")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(ip)
}
