package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go maFonction(&wg)
	wg.Wait()
	fmt.Println("I guess we're done here")
}

func maFonction(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("All done mate")
}
