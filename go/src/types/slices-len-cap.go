package main

import (
	"fmt"
)

/**
 * A slice has both a length and a capacity.
 * The length of a slice is the number of elements it contains.
 * The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
 * The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
 * You can extend a slice's length by re-slicing it, provided it has sufficient capacity. Try changing one of the slice operations in the example program to extend it beyond its capacity and see what happens.
 */

func printSlice(a []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
}

func main() {
	m := []int{1, 2, 34, 13, 55, 5}
	printSlice(m)

	m = m[:0]
	printSlice(m)

	m = m[:4]
	printSlice(m)

	m = m[2:]
	printSlice(m)
}
