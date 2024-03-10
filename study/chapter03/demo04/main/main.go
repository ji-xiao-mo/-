package main

import (
	"GOwork/study/chapter03/demo04/model"
	"fmt"
)

func main() {

	var i int = 10
	i = 50
	//i = 10.2不能改他的类型
	fmt.Println(i)

	//var i = 90  重复定义不行

	fmt.Println(model.HeroName)
}
