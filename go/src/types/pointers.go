package main

import "fmt"

/**
* & 获取内存地址
* * 获取内存地址的value
* The type *T is a pointer to a T value. Its zero value is nil.
* The & operator generates a pointer to its operand.
* The * operator denotes the pointer's underlying value.
 */
func main() {
	var h *int
	fmt.Println(h)
	i, j := 34, 78

	p := &i
	n := p
	m := &n

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(n)
	fmt.Println(*n)
	fmt.Println(m)
	fmt.Println(*m)

	*p = 32
	fmt.Println(i)
	p = &j
	*p = *p * 2
	fmt.Println(j)
}
