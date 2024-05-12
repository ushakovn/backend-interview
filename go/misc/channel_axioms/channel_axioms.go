package main

import (
	"flag"
	"log"
)

func main() {
	var ch chan struct{}

	opName := flag.String("operation-name", "", "")
	chType := flag.String("channel-type", "", "")
	chClose := flag.Bool("close-channel", false, "")

	flag.Parse()

	switch *chType {
	case "nil":
		ch = nil
	case "unbuffered":
		ch = make(chan struct{})
	case "buffered":
		ch = make(chan struct{}, 1)
	default:
		log.Fatalf("unexpected channel-type")
	}

	if *chClose {
		close(ch)
	}

	switch *opName {
	case "write":
		ch <- struct{}{}
	case "read":
		<-ch
	default:
		log.Fatalf("unexpected operation-name")
	}
}
