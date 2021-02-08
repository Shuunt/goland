package main

import (
	"flag"
	"fmt"
)

func main() {

	bFlag := flag.Bool("version", false, "numéro de version")
	flag.Parse()

	if *bFlag {
		fmt.Println(*bFlag)
		fmt.Println("http://url1.tld,http://url2.tld")
		return
	}

	//var time = ""
	//fmt.Printf("%v", time)
}
