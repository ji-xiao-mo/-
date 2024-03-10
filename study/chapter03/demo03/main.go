package main

import "fmt"

// 定义全局变量
var (
	n5    = 300
	n6    = 900
	name2 = "mary"
)

func main() {

	//演示go一次性声明多个变量(相同类型)
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//声明不同类型的变量
	var name, n4 = "tom", 100
	fmt.Println("name=", name, "n4=", n4)

	//输出全局变量
	fmt.Println(name2, n5, n6)
}
