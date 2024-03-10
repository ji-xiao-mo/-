package utils

var Age int
var Name string

//全局变量要初始化，在main.go中使用

func init() {
	Age = 18
	Name = "tom"
}
