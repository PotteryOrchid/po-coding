package main

import "fmt"

/**
* An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
 */
func main() {
	p := [3]int{2, 4, 6}

	var a []int = p[0:1]
	fmt.Println(a)
}
