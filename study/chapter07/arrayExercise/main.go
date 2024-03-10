package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//案例1
	var myChars [26]byte
	for i := 0; i < 26; i++ {
		myChars[i] = 'A' + byte(i)
	}
	for j := 0; j < 26; j++ {
		fmt.Printf("第%v个元素为%c\n", j+1, myChars[j])
	}

	//案例2
	//请求一个数组的最大值，得到相应的下标
	intarr := [4]int{1, 2, 3, 1}
	max := intarr[0]
	k := 0
	for i := 1; i < 4; i++ {
		if max < intarr[i] {
			max = intarr[i]
			k = i
		}
	}
	fmt.Println("最大值:", max, "下标", k)

	//案例3
	//随机生成5个数，反转打印
	//思路
	//1、随机生成用rand.Init()函数
	//2、得到随机数之后，放到数组中

	var arr03 [5]int
	for i := 0; i < 5; i++ {
		arr03[i] = rand.Intn(10)
	}
	fmt.Println(arr03)
	for j := 4; j >= 0; j-- {
		fmt.Println(arr03[j])
	}
}
