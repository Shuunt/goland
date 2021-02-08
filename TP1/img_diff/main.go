package main

import (
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
)

const PATH1 = "dummies/image_1.jpg"
const PATH2 = "dummies/image_2.jpg"
const PATH3 = "dummies/image_3.jpg"

func printContent(path string) int {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	h := sha256.New()
	write, err := h.Write(content)
	return write
}

func main() {
	var val1 = printContent(PATH1)
	var val2 = printContent(PATH2)
	var val3 = printContent(PATH3)

	if val1 != val2 && val1 != val3 {
		fmt.Println(PATH1)
	} else if val1 != val2 && val3 != val2 {
		fmt.Println(PATH2)
	} else {
		fmt.Println(PATH3)
	}
}
