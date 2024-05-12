package main

func merge(cs ...<-chan int) <-chan int {
	m := make(chan int)
	for _, cs := range cs {
		go func(c <-chan int) {
			for v := range c {
				m <- v
			}
		}(cs)
	}
	return m
}
