package main

import (
	"fmt"
)

// 一个函数test
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

// 一个函数
func getSum(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {

	//调用函数
	var n1 int = 10
	var n2 int = 13
	test(n1)
	fmt.Println("main() n1=", n1)

	sum, sub := getSum(n1, n2)
	fmt.Println("sum=", sum, "sub=", sub)

}
