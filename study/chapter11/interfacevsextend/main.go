package main

import (
	"fmt"
)

// Monkey结构体
type Monkey struct {
	Name string
}

func (m *Monkey) climbing() {
	fmt.Println(m.Name, "生来会爬树")
}

// 声明小鸟接口
type BirdAble interface {
	fly()
}

// LittleMonkey结构体
type LittleMonkey struct {
	Monkey
}

func (m *LittleMonkey) fly() {
	fmt.Println(m.Name, "学会飞行")
}

func main() {

	littlemonkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	littlemonkey.climbing()
	littlemonkey.fly()
}
