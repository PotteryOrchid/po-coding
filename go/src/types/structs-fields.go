package main

import (
	"fmt"
)

// Struct fields are accessed using a dot.
type Vertex struct {
	x int
	y int
}

func main() {
	m := Vertex{1, 2}
	fmt.Println(m.x)
}
