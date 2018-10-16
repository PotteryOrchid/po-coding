package main

import "fmt"

/**
 * It is common to append new elements to a slice, and so Go provides a built-in append function.
 * cap extend by 2^n.
 */
func main() {
	var a []int
	printSlice(a)

	a = append(a, 2)
	printSlice(a)

	a = append(a, 2, 3, 45)
	printSlice(a)

	a = append(a, 23, 33)
	printSlice(a)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
