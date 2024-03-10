package main

import (
	"fmt"
)

// 函数可以是一个数据类型，可以赋值给变量，通过该变量可以调用函数
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

// 函数可以当做形参传入
func myFun(han func(int, int) int, n1 int, n2 int) int {
	sel := han(n1, n2)
	return sel
}

// 加上案例，为函数类型取别名
type myFunType func(int, int) int

func myFun2(han myFunType, n1 int, n2 int) int {
	sel := han(n1, n2)
	return sel
}

// 支持对函数返回值命名
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sub = n1 - n2
	sum = n1 + n2
	return
}

func main() {
	//把函数名给一个变量，类似于起了一个新的名字
	a := getSum
	fmt.Printf("a的数据类型:%T \n", a)

	res1 := a(40, 33)
	fmt.Println("res1=", res1)

	var n1 = 1
	var n2 = 3

	res2 := myFun(getSum, n1, n2)
	fmt.Println("res2=", res2)

	//为int 取了一个别名
	type myInt int
	var num1 myInt
	var num2 int
	num1 = 1
	//定义了一个num2 int 但是不能直接把num1赋值给num2,因为类型不同
	num2 = int(num1) //要显式转类型
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	res3 := myFun2(getSum, n1, n2)
	fmt.Println("res3=", res3)

	//案例
	a1, b2 := getSumAndSub(1, 2)
	fmt.Println(a1, b2)
}
