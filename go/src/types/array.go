package main

import "fmt"

/**
* The type [n]T is an array of n values of type T.
* An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
 */
func main() {
	var v [2]string
	v[0] = "hello"
	v[1] = "world"

	i := [4]int{1, 3, 4, 5}

	j := []int{2, 45, 67}

	fmt.Println(v, v[1])
	fmt.Println(i)
	fmt.Println(j)
}
