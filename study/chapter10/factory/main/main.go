package main

import (
	"GOwork/study/chapter10/factory/model"
	"fmt"
)

func main() {
	// var s model.student
	// s.Name = "tom"
	// s.Score = 99.9
	// fmt.Println(s)

	//因为结构体首字母是小写，所以不能直接用结构体名定义
	//用返回值类型为结构体指针的函数调用
	var s = model.NewStudent("zhan", 22)
	fmt.Println(*s)
}
