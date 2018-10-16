package main

import "fmt"

type Ver struct {
	x, y int
}

func main() {
	var m = map[string]Ver{
		"java": Ver{33, 33},
		"go":   Ver{22, 22},
	}
	fmt.Println(m)

	n := map[string]Ver{
		"java":   {12, 12},
		"python": {23, 23},
	}
	fmt.Println(n)
}
