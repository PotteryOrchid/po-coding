package main

import "fmt"

/**
 * A slice literal is like an array literal without the length.
 */
func main() {
	a := []int{23, 34, 45, 2, 54}
	b := []bool{true, false, false}

	s := []struct {
		i int
		b bool
	}{{1, true}, {2, false}}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(s)
}
