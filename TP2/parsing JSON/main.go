package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	userName string `json:"userName"`
	Password string `json:"Password"`
}

func main() {
	//group := User{
	//	Login:    "Paul",
	//	Password: "pass123",
	//}
	filepath, err := os.Open("users.json")

	//Buffer
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.Encode(filepath)
	//b, err := json.Marshal(buffer)

	if err != nil {
		fmt.Println("error:", err)
	}
}
