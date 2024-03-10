package main

import (
	"GOwork/project/gin-friend/router"
)

func main() {
	//定义gin框架引擎，相关路由设置在文件router当中
	r := router.Router()

	//defer(延迟执行的函数（顺序类似于栈）) recover panic（让程序崩溃的函数）
	// defer fmt.Println(1)
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("捕获异常", err)
	// 	}
	// }()
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// panic("21")

	r.Run(":9999")
}
