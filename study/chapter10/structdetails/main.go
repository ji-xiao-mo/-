package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

// 结构体
type Rect struct {
	leftUp, rightDown Point //结构体类型的字段
}

type Rect02 struct {
	leftUp, rightDown *Point //结构体类型的字段
}

func main() {

	r1 := Rect{Point{1, 2}, Point{3, 4}}

	//r1在内存中有4个整数，在内存中连续分布
	fmt.Printf("r1.leftUp.x的地址是%p,r1.leftUp.y的地址是%p\n", &r1.leftUp.x, &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x的地址是%p,r1.rightDown.y的地址是%p\n", &r1.rightDown.x, &r1.rightDown.y)

	fmt.Println(r1)

	r2 := Rect02{&Point{10, 20}, &Point{30, 40}}
	fmt.Println(&r2.leftUp, &r2.rightDown)
	fmt.Println(r2.leftUp, r2.rightDown)

}
