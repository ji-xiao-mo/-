package main

import (
	"fmt"
)

// 编写一个学生考试系统
type Student struct {
	Name  string
	Age   int
	Score int
}

// 显示他的成绩
func (s *Student) showInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", s.Name, s.Age, s.Score)
}

// 设置成绩
func (s *Student) SetScor(score int) {
	if score < 0 || score > 100 {
		fmt.Println("设置有问题")
	} else {
		s.Score = score
	}
}

//小学生

type Pupil struct {
	Student
}

// 显示当前状态
func (p *Pupil) testing() {
	fmt.Println("小学生正在考试")
}

//大学生

type Graduate struct {
	Student
}

// 显示当前状态
func (g *Graduate) testing() {
	fmt.Println("大学生正在考试")
}

func main() {

	//当对结构体嵌入了匿名结构体，使用方法会发生变化
	pupil := &Pupil{}
	pupil.Name = "答案"
	pupil.Age = 11
	pupil.testing()
	pupil.SetScor(99)
	pupil.showInfo()

	gradu := &Graduate{}
	gradu.Name = "大学生"
	gradu.Age = 20
	gradu.testing()
	gradu.SetScor(88)
	gradu.showInfo()
}
