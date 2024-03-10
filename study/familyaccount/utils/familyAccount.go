package utils

import (
	"fmt"
)

type FamilyAccount struct {
	//用来接收输入的字符
	key string
	//声明一个变量用来控制循环
	loop bool
	//定义账户的余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义变量，记录是否有收支明细
	flag bool

	details string
}

// 编写工厂模式的构造方法，返回一个FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 0.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说 明",
	}

}

// 将显示明细变为方法
func (f *FamilyAccount) showDetails() {
	fmt.Println("--------------当前收支明细------------------")
	if f.flag {
		fmt.Println(f.details)
	} else {
		fmt.Println("当前没有收支")
	}
}

// 将登记收入变为方法
func (f *FamilyAccount) income() {
	fmt.Println("本次收入金额：")
	fmt.Scanln(&f.money)
	f.balance += f.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&f.note)
	//将收入情况拼接到details变量
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}

// 将登记支出变为方法
func (f *FamilyAccount) pay() {
	fmt.Println("本次支出金额：")
	fmt.Scanln(&f.money)
	if f.money > f.balance {
		fmt.Println("余额不足")
	} else {
		f.balance -= f.money
	}
	fmt.Println("本次收入说明：")
	fmt.Scanln(&f.note)
	//将收入情况拼接到details变量
	f.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", f.balance, f.money, f.note)
	f.flag = true
}

// 将退出系统变为方法
func (f *FamilyAccount) exit() {
	fmt.Println("你确认要退出吗? y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		} else {
			fmt.Println("输入有误,请重新输入")
		}
	}
	if choice == "y" {
		f.loop = false
	} else {
		fmt.Println("请输入正确选项")
	}

}

// 给该结构体绑定方法
// 显示主菜单
func (f *FamilyAccount) MainMenu() {
	for f.loop {
		fmt.Println("\n--------------家庭收支记账软件------------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Println("请选择(1-4)")

		fmt.Scanln(&f.key)

		switch f.key {
		case "1":
			f.showDetails()

		case "2":
			f.income()

		case "3":
			f.pay()

		case "4":
			f.exit()
		}
	}

}
