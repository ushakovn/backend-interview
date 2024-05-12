package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeStart := time.Now()

	<-worker()
	<-worker()

	fmt.Println(time.Since(timeStart).Seconds())
}

func worker() <-chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- rand.Int()
	}()

	return ch
}
