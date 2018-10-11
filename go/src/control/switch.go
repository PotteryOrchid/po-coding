package main

import (
	"fmt"
	"runtime"
)

/**
* Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
*/
func main() {
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Linux")
	case runtime.GOOS:
		x := 34 + 23
		if x > 23 {
			x = 54
		} else {
			x = 300
		}
		fmt.Println(x)
		fmt.Println("true")
	default:
		fmt.Println(os)
	}
}
