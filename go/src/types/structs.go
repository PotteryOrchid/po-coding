package main

import (
	"fmt"
	"reflect"
)

type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(Vertex{23, 45})
	m := Vertex{1, 2}
	fmt.Println(m)
	fmt.Println(reflect.ValueOf(&m).Elem().Field(1))
}
