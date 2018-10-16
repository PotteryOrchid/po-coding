package main

import "fmt"

var a = []int{2, 3, 4, 5}

func main() {
	for i, v := range a {
		fmt.Printf("%v %v\n", i, v)
	}
}
