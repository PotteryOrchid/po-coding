package main

import (
	"fmt"
	"math"
)

/**
* As a way to play with functions and loops, let's implement a square root function: given a number x, we want to find the number z for which z² is most nearly x (Newton`s method).
 */
func Sqrt(x float64) float64 {
	z := 1.0
	res := z*z - x
	for math.Abs(res) > 0.001 {
		res = z*z - x
		z -= res / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(3))
}
