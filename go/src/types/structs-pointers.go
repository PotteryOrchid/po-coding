package main

import (
	"fmt"
)

/**
* Struct fields can be accessed through a struct pointer.
* To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
* 注意：指针获取属性时需要将 *n 用小括号包含起来, 否则会报指针引用错误，因为 *n.Y 中 n.Y 是获取到的属性Y 为int类型，所以用指针的方法调用会报错。
 */
type Vertex struct {
	X int
	Y int
}

func main() {
	m := Vertex{1, 2}
	n := &m
	fmt.Println(n.X, (*n).Y)
}
