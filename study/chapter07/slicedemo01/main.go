package main

import (
	"fmt"
)

func main() {

	//演示切片的基本使用
	var intArr [5]int = [...]int{1, 22, 33, 66, 99}
	//定义切片
	//slice表示切片名
	//intArr[1:3] 表示引用数组intArr这个数组
	//引用数组的起始下标为1，最后下标为3（但不包含3）
	slice := intArr[1:3]
	fmt.Println("intArr=", intArr)         //[1 22 33 66 99]
	fmt.Println("slice的元素是", slice)        // [22 33]
	fmt.Println("slice的元素个数是", len(slice)) //2
	fmt.Println("slice的元素容量是", cap(slice)) //4

}
