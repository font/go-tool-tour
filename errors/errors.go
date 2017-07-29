package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

const delta = 0.00000000005

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}

	z := x
	z_prev := 0.0

	for i := 0; i < 10 && math.Abs(z_prev - z) > delta; i++ {
		z_prev = z
		z = z - (((z * z) - x) / (2 * z))
		fmt.Println(i, " = ", z)
	}

	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

