package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Login    string
	Password string
}

func main() {
	group := User{
		Login:    "Paul",
		Password: "pass123",
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
