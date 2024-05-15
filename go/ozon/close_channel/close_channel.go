package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // nit: fixed
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
