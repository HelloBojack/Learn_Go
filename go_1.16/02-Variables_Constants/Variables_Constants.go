package main

import "fmt"

// 全局变量
var a = 1
var (
	b = 2
)
var c, d = 3, 4

// c:=3

const f = 123

func main() {
	e := 5
	var g int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(g)
}
