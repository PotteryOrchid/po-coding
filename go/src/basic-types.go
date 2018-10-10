package main

import (
	"fmt"
	"math/cmplx"
)

/**
* The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems. When you need an integer value you should use int unless you have a specific reason to use a sized or unsigned integer type.
*/
var (
	ToBe bool = true
	MaxInt uint64 = 1<<64 - 1
	c complex128 = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Printf("Type: %T value: %v\n",ToBe,ToBe)
	fmt.Printf("Type: %T value: %v\n",MaxInt,MaxInt)
	fmt.Printf("Type: %T value: %v\n",c,c)
}
