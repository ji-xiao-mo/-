package main

import (
	"fmt"
	"math"
)

func main() {
	var score int32
	fmt.Println("请输入考试成绩")
	fmt.Scanln(&score)

	//练习一、多分支判断

	if score == 100 {
		fmt.Println("奖励一辆车")
	} else if score > 80 && score <= 99 {
		fmt.Println("奖励一台三星")
	} else if score >= 60 && score <= 80 {
		fmt.Println("奖励一个鼠标")
	} else {
		fmt.Println("奖励无")
	}

	//练习二、求ax2+bx+c=0方程的根
	var a float64
	var b float64
	var c float64
	var x1 float64
	var x2 float64
	fmt.Println("请输入参数 a b c")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Println(a)

	var m float64 = b*b - 4*a*c

	if m > 0 {
		x1 = (-b + math.Sqrt(m)) / 2 * a
		x2 = (-b - math.Sqrt(m)) / 2 * a
		fmt.Printf("x1=%v x2=%v", x1, x2)

	} else if m == 0 {
		x1 = (-b + math.Sqrt(m)) / 2 * a
		fmt.Println(x1)
	} else {
		fmt.Println("无解")
	}

}
