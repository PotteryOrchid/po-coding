package main

import "fmt"

func fibonacci() func() int {
	fib := 0
	return func() int {
		f, s, i := 0, 1, 0
		for i < fib {
			f, s = s, f+s
			i++
		}
		fib++
		return f

	}
}

func main() {
	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
