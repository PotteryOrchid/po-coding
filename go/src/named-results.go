package main

import "fmt"

func split(x int) (a, b int) {
	a = x * 4 / 9
	b = x - a
	return
}

func main(){
	fmt.Println(split(18))
}
