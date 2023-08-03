package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("This happens asynchronously")
		wg.Done()
	}()    //These parentheses at the end invoke the function

	fmt.Println("This is synchronous")
	wg.Wait() // called after all the concurrent work has been registered
}