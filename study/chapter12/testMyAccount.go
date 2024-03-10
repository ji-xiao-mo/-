package main

import (
	"fmt"
)

func main() {
	//用来接收输入的字符
	key := ""

	//声明一个变量用来控制循环
	loop := true

	//定义账户的余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义变量，记录是否有收支明细
	flag := false

	details := "收支\t账户金额\t收支金额\t说 明"

	//显示主菜单
	for loop {
		fmt.Println("\n--------------家庭收支记账软件------------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Println("请选择(1-4)")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("--------------当前收支明细------------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支")
			}

		case "2":
			fmt.Println("本次收入金额：")
			fmt.Scanln(&money)
			balance += money
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将收入情况拼接到details变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true

		case "3":
			fmt.Println("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足")
				break
			} else {
				balance -= money
			}
			fmt.Println("本次收入说明：")
			fmt.Scanln(&note)
			//将收入情况拼接到details变量
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true

		case "4":
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
				loop = false
			}

		default:
			fmt.Println("请输入正确选项")

		}

	}

}
