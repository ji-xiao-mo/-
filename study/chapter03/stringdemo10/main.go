package main

import (
	"fmt"
)

func main() {
	var address string = "北京长城 hello world"
	fmt.Println(address)

	var str3 string = `//有符号整数
	var i int8 = 127
	//var i int8 = 128  超出范围，报错
	fmt.Println("i=", i)

	//无符号整数uint 范围055
	var k uint8 = 255
	`
	fmt.Println(str3)
	//字符串拼接方式

	var str = "fwef" + "ger"
	fmt.Println(str)
	//当拼接非常的长的时候，加法要分行
	var str1 = "fwe" + "wfge" +
		"iohjoi"
	fmt.Println(str1, "\n")

	var a int
	var b float32
	var c float64
	var d bool
	var e string

	fmt.Println(a, b, c, d, e)
}
