package main

import "fmt"

/**
* Constants are declared like variables, but with the const keyword.
* Constants can be character, string, boolean, or numeric values.
* Constants cannot be declared using the := syntax.
*/
const Pi = 3.14

func main(){
	const Msg = "hello world"
	fmt.Println(Msg,Pi)
}
