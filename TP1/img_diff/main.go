package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

func printContent(path string) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println(content)

	h := sha256.New()

	fmt.Println(h.Write(content))
}

func main() {
	printContent("dummies/image_1.jpg")
	//printContent("dummies/image_2.jpg")
	//printContent("dummies/image_3.jpg")
}
