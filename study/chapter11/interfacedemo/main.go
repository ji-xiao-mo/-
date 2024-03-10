package main

import (
	"fmt"
)

// 声明定义一个借口
type Usb interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Usb2 interface {
	//声明两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
}

// 让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println("手机停止工作")
}

// 让相机实现功能
type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作")
}

// 让计算机实现功能
type Computer struct {
}

// 一个接口类型的，定义了方法未定义，
func (c Computer) Working(usb Usb2) {

	//通过usb接口变量来调用Start和Stop函数
	usb.Start()
	usb.Stop()
}

func main() {

	//测试
	//先创建结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	//关键点
	computer.Working(phone)
	computer.Working(camera)

}
