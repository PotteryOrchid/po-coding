package main

import (
	"fmt"
)

/**
* When slicing, you may omit the high or low bounds to use their defaults instead.
* The default is zero for the low bound and the length of the slice for the high bound.
 */
func main() {
	m := []int{1, 2, 34, 13, 55, 5}
	a := m[2:5]
	b := m[:]
	d := m[:4]
	fmt.Println(a, b, d)
}
