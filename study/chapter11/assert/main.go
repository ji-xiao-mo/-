package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	//空接口赋值给一个变量要用类型断言
	var b Point = a.(Point)
	fmt.Println(b)
}
