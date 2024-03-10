package main

import (
	"fmt"
	"os"
)

func main() {

	//使用ioutil.ReadFile一次性读取文件
	file := "/Users/hanzongzheng/Desktop/GOwork/chapter14/filedemo02/test.txt"
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Read this file :", err)
	}
	fmt.Println(string(content)) //输出的是[]byte

}
