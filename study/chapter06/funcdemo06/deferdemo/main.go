package main

import (
	"fmt"
)

// 遇到defer之后，语句会压入defer栈中，然后先入后出
func sum(n1 int, n2 int) int {
	defer fmt.Println(n1) //3.10  先入的所以后出,压入栈的时候，n1被同时压入栈
	defer fmt.Println(n2) //2.30  后入的所以先出，n2被同时压入栈
	//增加一句话
	n1++             //n1 = 11
	n2++             //n2 = 31
	res := n1 + n2   // res = 42
	fmt.Println(res) //1. 40
	return res
}

func main() {

	res := sum(10, 30)
	fmt.Println("res=", res) //4. res= 40

}
