package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getImg(args []string) *string {
	var listUrl = strings.Split(args[0], ",")
	for _, element := range listUrl {
		start := time.Now()
		var img, err = http.Get(element)
		if err != nil {
			log.Fatal(err)
		}
		//var finalImg ,_ = ioutil.ReadAll(img.Body)
		//fmt.Println(finalImg)
		//fmt.Println(img)
		//fmt.Println(listUrl)
		defer img.Body.Close()
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("%v\n", elapsed)
	}
	fmt.Println(listUrl)
	return nil
}

func main() {
	if err := getImg(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
