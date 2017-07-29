package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	n1, n2, fib := 1, 0, 0

	return func() int {
		if fib == 0 {
			fib = n1 + n2
			return 0
		}

		n := fib
		fib = n1 + n2
		n1, n2 = fib, n1
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

