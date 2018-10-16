package main

import (
	"fmt"
	"strings"
)

/**
 * Slices can contain any type, including other slices.
 */
func main() {
	a := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	a[0][0] = "X"
	a[0][1] = "X"
	a[1][0] = "X"
	a[1][2] = "0"
	a[2][1] = "0"

	for i := 0; i < len(a); i++ {
		fmt.Printf("%s\n", strings.Join(a[i], " "))
	}

	fmt.Println(strings.Join([]string{"sds", "sds"}, "++"))
}
