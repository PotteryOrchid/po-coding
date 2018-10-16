package main

import "fmt"

type Ver struct {
	x, y int
}

// map must be init by make func.
func main() {
	var m map[string]Ver
	m = make(map[string]Ver)
	m["hello"] = Ver{12, 13}
	fmt.Println(m)
	fmt.Println(m["hello"])

	n := make(map[string]Ver)
	n["world"] = Ver{99, 99}
	fmt.Println(n["world"])

	n["kid"] = Ver{88, 88}
	fmt.Println(n)
}
