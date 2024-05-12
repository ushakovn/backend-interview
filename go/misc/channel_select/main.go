package main

import (
	"fmt"
)

func main() {
	msgCh := make(chan string, 1)
	close(msgCh)

	doneCh := make(chan struct{})
	close(doneCh)

	select {
	case msgCh <- "":
		fmt.Println("msg sent")
	case <-doneCh:
		fmt.Println("stop signal received")
	}
}
