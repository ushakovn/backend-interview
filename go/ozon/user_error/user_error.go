package main

func main() {
	println(handle().Error())
}

func handle() error {
	return &err{
		s: "something went wrong...",
	}
}

type err struct {
	s string
}

func (e err) Error() string {
	return e.s
}
