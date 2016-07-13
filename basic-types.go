package main

import (
	"fmt"
	"math/cmplx"
)

var (
	// Defined Bool
	ToBe bool = false
	// Defined uint64 type
	MaxInt uint64 = 1<<64 - 1
	// Defined Complex and calculated sqrt using cmplx
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// style in which you want the data output
	const f = "%T(%v)\n"
	// using fmt printed out the value the way we wanted
	// Printed ToBe, ToBe 
	fmt.Printf(f, ToBe, ToBe)
	// Printed MaxInt of MaxInt
	fmt.Printf(f, MaxInt, MaxInt)
	// Printed output of z complex i.e calcualted complex128 type using cmplx.Sqrt function
	fmt.Printf(f, z, z)
}