package main

import (
	"fmt"
)

func main() {
	var num int = 9

	fmt.Println("num的地址=", &num)

	var ptr *int = &num
	*ptr = 10
	fmt.Println("num的值=", num)

}
