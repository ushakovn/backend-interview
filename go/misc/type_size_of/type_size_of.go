package main

import "unsafe"

func main() {
	var (
		a     interface{}
		s     struct{}
		r     string
		b     bool
		i     int
		i8    int8
		i32   int32
		i64   int64
		f32   float32
		f64   float64
		runes []rune
		slice []interface{}
		array [0]interface{}
	)
	println("size of default go types in bytes")
	println("interface{}:\t", unsafe.Sizeof(a))
	println("struct{}:\t", unsafe.Sizeof(s))
	println("string:\t\t", unsafe.Sizeof(r))
	println("bool:\t\t", unsafe.Sizeof(b))
	println("int:\t\t", unsafe.Sizeof(i))
	println("int8:\t\t", unsafe.Sizeof(i8))
	println("int32:\t\t", unsafe.Sizeof(i32))
	println("int64:\t\t", unsafe.Sizeof(i64))
	println("float32:\t", unsafe.Sizeof(f32))
	println("float64:\t", unsafe.Sizeof(f64))
	println("[]rune:\t\t", unsafe.Sizeof(runes))
	println("[]interface{}:\t", unsafe.Sizeof(slice))
	println("[0]interface{}:\t", unsafe.Sizeof(array))
}

// stdout
//
//size of default go types in bytes
//interface{}:     16
//struct{}:        0
//string:          16
//bool:            1
//int:             8
//int8:            1
//int32:           4
//int64:           8
//float32:         4
//float64:         8
//[]rune:          24
//[]interface{}:   24
//[0]interface{}:  0

// string size explanation:
//
// type StringHeader struct {
//        Data uintptr
//        Len  int
//}
