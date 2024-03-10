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

type Phone struct {
	Name string
}

// 让Phone 实现 Usb接口的方法
func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

// Phone独有的方法
func (p Phone) Call() {
	fmt.Println("手机打电话")
}

// 让相机实现功能
type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println(c.Name, "相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "相机停止工作")
}

// 让计算机实现功能
type Computer struct {
}

// 一个接口类型的，定义了方法未定义，
func (c Computer) Working(usb Usb) {

	//通过usb接口变量来调用Start和Stop函数
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用其本身call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {

	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"oppo"}
	usbArr[2] = Camera{"索尼"}

	var computer Computer

	for _, v := range usbArr {
		computer.Working(v)
	}

	fmt.Println(usbArr)

}
