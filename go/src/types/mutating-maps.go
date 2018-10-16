package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["first"] = 12
	fmt.Println("value is ", m["first"])

	m["first"] = 23
	fmt.Println("value is ", m["first"])

	delete(m, "first")
	fmt.Println("value is ", m["first"])

	v, flag := m["first"]
	fmt.Println("Value is ", v, ", Is present? ", flag)

}
