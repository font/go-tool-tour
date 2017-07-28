package main

import (
	"fmt"
	"math"
)

const delta = 0.00000000005

func Sqrt(x float64) float64 {
	z := x
	z_prev := 0.0
	for i := 0; i < 10 && math.Abs(z_prev - z) > delta; i++ {
		z_prev = z
		z = z - (((z * z) - x) / (2 * z))
		fmt.Println(i, " = ", z)
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

