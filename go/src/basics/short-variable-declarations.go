package main

import "fmt"

var x,y int = 23,43

/**
* Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
* Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/
func main(){
	c, python, java := true,false,"hello"
	fmt.Println(x,y,c,python,java)
}
