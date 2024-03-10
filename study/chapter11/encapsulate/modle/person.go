package modle

import "fmt"

type person struct {
	Name string
	age  int
	sal  float64
}

// 现在其他包创建不了person类型的包，需要工厂模式的函数
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// 为了访问age和sal，构建set函数可以为其设置值，get方法可以访问其值
// 设置年龄之前必须要对设置的信息进行判断是否合规
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) SetSal(sal float64) {
	if sal >= 3000 {
		p.sal = sal
	} else {
		fmt.Println("给的不够啊🤔")
	}
}

func (p *person) GetPerson() (int, float64) {
	return p.age, p.sal
}
