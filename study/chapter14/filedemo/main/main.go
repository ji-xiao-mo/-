package main

import (
	"fmt"
	"os"
)

func main() {

	//打开文件
	file, err := os.Open("/Users/hanzongzheng/Desktop/GOwork/test.go")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}

	//输出文件
	fmt.Printf("file=%v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("关闭错误")
	}
}
