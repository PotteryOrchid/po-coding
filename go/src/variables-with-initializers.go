package main

import "fmt"

/**
* A var declaration can include initializers, one per variable.
* If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/
var x,y int = 23,43

func main(){
	var c, python, java = true,false,"hello"
	fmt.Println(x,y,c,python,java)
}
