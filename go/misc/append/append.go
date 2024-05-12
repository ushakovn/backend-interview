package main

import "fmt"

func main() {
  source := []int{1, 2, 3}
  log("source init", source)

  trimmed := source[1:2]
  log("trimmed after trim", trimmed)

  appended := append(trimmed, 4)
  log("appended after append", appended)
  log("source after append", source)
}

func log(op string, s []int) {
  fmt.Printf("%s\n", op)
  fmt.Printf("value = %v\n", s)
  fmt.Printf("len = %d\n", len(s))
  fmt.Printf("cap = %d\n\n", cap(s))
}
