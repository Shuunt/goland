package main

import "fmt"
import "flag"

const VERSION = "1.0"

func main() {

	bFlag := flag.Bool("version", false, "numéro de version")
	flag.Parse()

	if *bFlag {
		fmt.Println(*bFlag)
		fmt.Println(VERSION)
		return
	}
}
