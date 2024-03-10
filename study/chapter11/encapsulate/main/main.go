package main

import (
	"GOwork/study/chapter11/encapsulate/modle"
	"fmt"
)

func main() {
	p := modle.NewPerson("jack")
	p.SetAge(22)
	p.SetSal(999999)
	age, sal := p.GetPerson()
	fmt.Printf("%v的信息:", p.Name)
	fmt.Println("年龄为", age, "薪水为", sal)
}
