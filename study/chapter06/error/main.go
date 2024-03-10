package main

import (
	"errors"
	"fmt"
)

func test() {
	//使用defer + recover 来捕获和处理异常
	defer func() {
		err := recover() //内置函数，可以捕获到异常
		if err != nil {  //说明捕获到异常
			fmt.Println("err=", err)
			//这里可以将错误信息发送给管理员
			fmt.Println("发送邮件给管理员")
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println(res)
}

// 函数用于读取配置文件
// 如果传入文件名错误，则返回自定义错误
func readConf(name string) (err error) {
	if name == "config.ini" {
		//读取该文件
		return nil
	} else {
		//返回一个自定义的错误
		return errors.New("读取文件错误")
	}
}

func test02() {
	err := readConf("config.in")
	if err != nil {
		//输出该错误，终止程序
		panic(err)
	}
	fmt.Println("ffwefew")
}

func main() {

	// //测试
	// test()

	// fmt.Println("错误")

	//测试自定义错误
	test02()
}
