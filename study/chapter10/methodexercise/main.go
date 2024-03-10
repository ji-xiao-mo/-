package main

import (
	"fmt"
)

//声明一个结构体Circle, 字段为radius
//声明一个方法area和Circle绑定，可以返回面积
//画出area执行过程+说明

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	res := c.radius * c.radius * 3.14
	return res
}

func main() {

	var cir Circle
	cir.radius = 4.0
	ar := cir.area()
	fmt.Println("面积是=", ar)

}
