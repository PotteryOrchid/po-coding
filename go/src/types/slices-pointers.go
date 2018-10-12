package main

import "fmt"

/**
 * A slice does not store any data, it just describes a section of an underlying array.
 * Changing the elements of a slice modifies the corresponding elements of its underlying array.
 * Other slices that share the same underlying array will see those changes.
 */
func main() {
	p := [6]string{"hello", "world", "I", "am", "java", "language"}

	var a []string = p[2:5]
	b := p[2:6]
	fmt.Println(a)
	fmt.Println(b)

	b[2] = "go"
	fmt.Println(p)
	fmt.Println(a)
	fmt.Println(b)
}
