package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("original print : %v: %v\n", name, ip)
		fmt.Printf("print with (a) version : %v: %v\n", name, fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3]))
		stringFormated := []byte("")
		for ind, ipNb := range ip {
			stringFormated = strconv.AppendInt(stringFormated, int64(ipNb), 10)
			if ind < 3 {
				stringFormated = append(stringFormated, 46)
			}
		}
		fmt.Printf("print with (b) version : %v: %s\n", name, stringFormated)
	}
}
