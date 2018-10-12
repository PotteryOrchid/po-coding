package main

import (
	"fmt"
)

/**
 * A struct literal denotes a newly allocated struct value by listing the values of its fields.
 * You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
 * The special prefix & returns a pointer to the struct value.
 */
type Vertex struct {
	X int
	Y int
}

func main() {
	m := Vertex{1, 2}
	n := Vertex{X: 23}
	o := Vertex{}
	p := &Vertex{1, 2}
	fmt.Println(m, n, o, p)
}
