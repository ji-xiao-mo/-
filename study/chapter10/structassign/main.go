package main

import (
	"fmt"
)

type Stu struct {
	Name string
	Age  int
}

func main() {

	//在创建结构体变量时，直接指定字段的值
	//方式一
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明～", 20}

	//方式二
	//创建结构体变量时，把字段名和字段值直接写在一起,不必要求赋值顺序和定义相同
	var stu3 = Stu{
		Name: "jack",
		Age:  25,
	}

	fmt.Println(stu1, stu2, stu3)
}
