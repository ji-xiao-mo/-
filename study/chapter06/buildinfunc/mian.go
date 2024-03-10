package main

import (
	"fmt"
)

func main() {

	num1 := 100
	fmt.Printf("num1的类型是%T, num1的值=%v, num1的地址%v\n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2的类型是%T, num2的值=%v, num2的地址%v\n", num2, num2, &num2)
	fmt.Println(*num2)

}
