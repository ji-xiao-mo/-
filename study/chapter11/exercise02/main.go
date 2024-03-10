package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1、声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

// 2、声明一个Hero结构体切片类型
type HeroSlice []Hero

// 3、实现InterSlice接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less方法是决定按照什么标准来进行排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}

func (hs HeroSlice) Swap(i, j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
	// 等价一句话  hs[i], hs[j] = hs[j], hs[i]
}

func main() {

	//测试
	var heroes HeroSlice //前面把[]hero重命名为HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%v", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//循环创建切片，将hero append到heroes结构体切片当中去
		heroes = append(heroes, hero)
	}

	//看排序前的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}

	//调用sort.Sort
	sort.Sort(heroes)
	fmt.Println("----------排序之后----------")

	//看排序后的顺序
	for _, v := range heroes {
		fmt.Println(v)
	}
}
