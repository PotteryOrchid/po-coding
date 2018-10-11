package main

import "fmt"

//The init and post statements are optional.
func main() {
	// sum init value must is not 0 (0 代表参数未初始化)
	sum := 2
	for sum < 1008 {
		sum += sum
	}
	fmt.Println(sum)
}
