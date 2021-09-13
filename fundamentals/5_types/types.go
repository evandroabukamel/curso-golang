package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// numerics
	fmt.Println(1, 2, 1000)
	fmt.Println("Integer literal is", reflect.TypeOf(32000))
	fmt.Println("Integer literal is", reflect.TypeOf(-32000))

	// Unsigned int are: uint8, uint16, uint32, uint64
	var b byte = 255
	fmt.Println("\nbyte b literal is", reflect.TypeOf(b))

	// Signed int are: int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("\nMaximun value of int is", i1)
	fmt.Println("i1 literal is", reflect.TypeOf(i1))

	// Represents a mapped value from Unicode table (int32)
	var i2 rune = 'a'
	fmt.Println("\nrune literal is", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Real numbers are: float32, float64
	var x1 float32 = 49.99
	var x2 = float32(49.99)
	fmt.Println("\nx1 literal is", reflect.TypeOf(x1))
	fmt.Println("x2 literal is", reflect.TypeOf(x2))
	// float64
	fmt.Println("float literal is", reflect.TypeOf(49.99))

	// Boolean
	bo := true
	fmt.Println("\nbo literal is", reflect.TypeOf(bo))
	fmt.Println("The oposite of bo is", !bo)

	// String
	s1 := "Hi, my name is Evandro"
	fmt.Println("\n" + s1 + "!")
	fmt.Println("s1 length is", len(s1))

	// Multiline string
	s2 := `Hello,
	my name is
	Evandro`
	fmt.Println(s2 + "!")
	fmt.Println("s2 length is", len(s2))

	// Char = int32
	char := 'a'
	fmt.Println("\nchar literal is", reflect.TypeOf(char))
	fmt.Println(char)
}
