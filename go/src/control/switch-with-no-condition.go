package main

import (
	"fmt"
	"time"
)

/**
* Switch without a condition is the same as switch true.
 */
func main() {
	fmt.Println("Go runs on")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	default:
		fmt.Println("evening")
	}
}
