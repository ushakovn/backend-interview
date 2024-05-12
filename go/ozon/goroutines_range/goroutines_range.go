package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4, 5}
	for _, val := range values {
		go func(v int) { // nit: fixed
			fmt.Println(v)
		}(val)
	}
	time.Sleep(100 * time.Millisecond)
}
