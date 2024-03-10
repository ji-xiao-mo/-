package main

import (
	"fmt"
)

func main() {
	var j int
	for i := 1; i < 100; i++ {
		if i%9 == 0 {
			fmt.Println(i)
			j += 1
		}
	}
	fmt.Println(j)
}
