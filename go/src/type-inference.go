package main

import "fmt"

/**
* When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
*/
func main(){
	var i = 32
	x := 'x'
	fmt.Printf("i type is %T and x type is %T\n",i,x)
}
