package main

import (
	"fmt"
)

// 定义一个struct结构体，将cat的各个字段属性，放入cat结构体中
type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1
	var p1 Person

	//方式2
	p2 := Person{"mary", 20}
	fmt.Println(p1)
	fmt.Println(p2)

	//方式3
	var person *Person = &Person{}
	(*person).Name = "fwe"
	//上面的是指针，但是使用起来比较别扭，设计者为了方便使用
	//支持下面这样直接像方式1或2的方式来使用
	person.Age = 55
	fmt.Println(*person)

}
