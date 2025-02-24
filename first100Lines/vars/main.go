package main

import "fmt"

/* Declare a single variable */
var a int

/* Declare a group of variables */
var (
	b bool
	c float32
	d string
)

func main() {
	a = 42
	b, c = true, 32.0
	d = "string"
	fmt.Println(a, b, c, d)

	// short-handed define variables without declare
	e := 42
	f, g := true, 32.0
	h := "string"
	fmt.Println(e, f, g, h)

	// ----------Types------------
	// User-specified types
	const aa int32 = 12        //32-bit integer
	const bb float32 = 20.5    //32-bit float
	var cc complex128 = 1 + 4i //128-bit complex number
	var dd uint16 = 14         // 16-bit unsigned integer

	// Default types
	n := 42              // int
	pi := 3.14           // float64
	t, f := true, false  // bool
	s := "Go is awesome" // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", aa, bb, cc, dd)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, t, f, s)
}
