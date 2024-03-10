package main

import (
	"fmt"
)

func main() {

	//创建一个可以存放3个int类型的管道
	var intChan chan int = make(chan int, 3)

	//intChan类型是引用类型
	fmt.Printf("intChan 的值=%v intChan 本身的地址=%p\n", intChan, &intChan)

	//向管道写入数据
	intChan <- 10
	num := 211
	intChan <- num

	//注意点，当给管道写入数据时，不能超过其容量
	// intChan <- 10
	// intChan <- 10

	//看管道的长度和容量
	fmt.Printf("channel len = %v cap = %v\n", len(intChan), cap(intChan))

	//从管道中读取数据
	num2 := <-intChan
	fmt.Println(num2)
	fmt.Printf("channel len = %v cap = %v\n", len(intChan), cap(intChan))

	//在没有使用协程的情况下，如果管道的数据全部取出来，再取就会报告错误
	num3 := <-intChan

	fmt.Println(num3)
	fmt.Printf("channel len = %v cap = %v\n", len(intChan), cap(intChan))

}
