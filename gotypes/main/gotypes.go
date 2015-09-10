package main

import ("fmt"
		"math/cmplx")

var (
	goIsFun bool = true
	maxInt uint64 = 1 <<64 -1
	complex complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const formatter  = "%T(%v) \n"
	fmt.Printf(formatter,goIsFun,goIsFun)
	fmt.Printf(formatter,maxInt,maxInt)
	fmt.Printf(formatter,complex,complex)
}