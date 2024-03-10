package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 1)给Person结构体添加speak方法，输出 xxx是一个好人
func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

// 2)给Person结构体添加jisuan方法，可以计算1+...1000的结果
// 说明方法体内和函数一样，可以进行各种运算
func (p Person) jisuan(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

// 给Person结构体添加getSum方法，可以计算两个数的和，并返回结果
func (p Person) getSum(n1 int, n2 int) {
	fmt.Println(n1 + n2)
}

// 给Person类型绑定一份方法
func (p Person) test() {
	fmt.Println("test()=", p.Name)
}

func main() {

	var p Person
	p.Name = "tom"
	p.test() //调用方法，把p传到方法中去
	p.speak()
	fmt.Println(p.jisuan(10))
	p.getSum(5, 0)

}
