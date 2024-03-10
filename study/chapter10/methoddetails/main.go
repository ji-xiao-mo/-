package main

import (
	"fmt"
)

type integer int

func (i integer) print() {
	i = 3
	fmt.Println("i=", i)
}

func (i *integer) print02() {
	*i = 55
}

func main() {

	var i integer = 99
	i.print() //把main栈中的值拷贝到print02栈当中，
	//函数当中改变i值不影响main栈中的i值
	fmt.Println(i)
	i.print02() //是地址拷贝，直接改变值
	fmt.Println(i)

}
