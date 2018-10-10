package main

import "fmt"

func getRes(x,y string) (string, string) {
	return x, y
}

func main() {
	//a,b := getRes("hello","monky")
	fmt.Println(getRes("hello","kitty"))
}
