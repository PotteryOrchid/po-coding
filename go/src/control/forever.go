package main

import "fmt"

//If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
func main() {
	fmt.Println("start")
	for {
	}
	fmt.Println("end")
}
