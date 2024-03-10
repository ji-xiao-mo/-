package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("/Users/hanzongzheng/Desktop/GOwork/chapter14/filedemo02/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)

	}

	//当函数退出时,要及时关闭file
	defer file.Close() //等到函数结束时执行

	//创建一个*Reader ,带缓冲
	reader := bufio.NewReader(file)
	//循环读取文件内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
}
