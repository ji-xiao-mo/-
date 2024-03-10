package main

import (
	"fmt"
)

type student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (stu *student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
		stu.name, stu.gender, stu.age, stu.id, stu.score)
	return infoStr
}

// 创建Box结构体 声明长宽高，
type Box struct {
	len   float64
	width float64
	high  float64
}

func (b *Box) tiji() float64 {
	res := b.high * b.len * b.width
	return res
}

// 景区门票案例
type Visitor struct {
	Name string
	Age  int
}

func (v Visitor) jiage() {
	if v.Age > 18 {
		fmt.Printf("%v的年龄为:%v", v.Name, v.Age)
		fmt.Println("应收取门票20")
	} else {
		fmt.Println("免费")
	}
}

func main() {
	var st = student{
		"tom", "男", 32, 3, 89,
	}
	fmt.Println(st.say())

	//长方体体积测试代码
	var B Box
	fmt.Println("请输入长方体的长：")
	fmt.Scanln(&B.len)
	fmt.Println("请输入长方体的宽：")
	fmt.Scanln(&B.width)
	fmt.Println("请输入长方体的高：")
	fmt.Scanln(&B.high)
	fmt.Println("长方体的体积为：", B.tiji())

	//景区门票案例
	var V Visitor
	fmt.Println("请输入姓名")
	fmt.Scanln(&V.Name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&V.Age)
	V.jiage()
}
