package main

func main() {
	const (
		numInt   int64   = 100
		numFloat float64 = 100.00
	)
	// compilation error: invalid operation: division by zero
	// println(numInt / 0)

	println(divideByZero(numFloat)) // println: +Inf
	println(divideByZero(numInt))   // panic: runtime error: integer divide by zero

}

func divideByZero[T int64 | float64](num T) T {
	return num / 0
}
