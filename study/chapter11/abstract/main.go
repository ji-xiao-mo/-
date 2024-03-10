package main

import (
	"fmt"
)

// 定义几个结构体
type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 方法
// 1、存款
func (account *Account) Deposite(money float64, pwd string) {
	//核对密码
	if pwd != account.Pwd {
		fmt.Println("密码错误")
		return
	}
	//判断金额
	if money <= 0 {
		fmt.Println("输入的金额错误")
		return
	}

	account.Balance += money
	fmt.Println("存款成功")
}

func main() {

	var account Account
	account.Balance = 22
	account.Pwd = "12203758"
	account.Deposite(20, "12203758")
	fmt.Println(account.Balance)

}
