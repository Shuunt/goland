package main

import (
	"fmt"
	"strconv"
	"time"
)

type IPAddr [4]byte

type MyError struct {
	When time.Time
	What string
}

func (err MyError) Error() string {
	return fmt.Sprintf("Error at: %v this happened: %s", err.When, err.What)
}

func run() error {
	return MyError{
		When: time.Now(),
		What: "test",
	}
}

type Inter interface {
}

func PrintIt(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Println("The type is int !")
	case string:
		fmt.Println("The type is string !")
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

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

	PrintIt(int(1))

	err := run()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("no error")

}
