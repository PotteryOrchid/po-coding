package main

import "fmt"

func main() {
	var a = make([]int, 5)

	for i := range a {
		a[i] = 1 << uint(i)
	}

	for _, v := range a {
		fmt.Printf("%v\n", v)
	}
}
