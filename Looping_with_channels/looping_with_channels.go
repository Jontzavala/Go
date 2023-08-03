package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)   // allows us to stop the loop so the function isn't waiting on the 11 number
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}