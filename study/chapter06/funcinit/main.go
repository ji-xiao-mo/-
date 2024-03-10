package main

import (
	"GOwork/study/chapter06/funcinit/utils"
	"fmt"
)

func test() int {
	fmt.Println("test()")
	return 90
}

// 先执行全局变量
var age = test() //1
// 完成初始化工作
func init() {
	fmt.Println("init()") //2
}

func main() {
	fmt.Println("main()", age) //3
	fmt.Println(utils.Age, utils.Name)
}
