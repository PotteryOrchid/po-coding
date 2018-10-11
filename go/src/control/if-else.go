package main

import (
	"fmt"
	"math"
)

/**
* Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.
* Like for, the if statement can start with a short statement to execute before the condition.
* Variables declared by the statement are only in scope until the end of the if.
 */
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
