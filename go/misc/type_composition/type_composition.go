package main

type Unpack interface {
	Unpack() string
}

type Bundle struct {
	Box
}

func NewBundle() Unpack {
	return new(Bundle)
}

func (c *Bundle) Unpack() string {
	return "bundle unpacking"
}

type Box struct{}

func NewBox() Unpack {
	return new(Box)
}

func (b *Box) Unpack() string {
	return "box unpacking"
}

func main() {
	unpack := NewBundle()

	_, ok := unpack.(*Bundle)
	println(ok)

	_, ok = unpack.(*Box)
	println(ok)

	unpack = NewBox()

	_, ok = unpack.(*Bundle)
	println(ok)

	_, ok = unpack.(*Box)
	println(ok)
}
