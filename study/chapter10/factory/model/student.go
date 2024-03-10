package model

// 定义一个结构体
type student struct {
	Name  string
	Score float64
}

//因为student结构体首字母是小写，因此只能在model使用
//通过工厂模式来解决

func NewStudent(n string, s float64) *student {
	//返回值是指针类型，所以存的应该是地址
	return &student{
		Name:  n,
		Score: s,
	}
}
