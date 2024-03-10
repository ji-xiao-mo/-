package main

import (
	"fmt"
)

func swap(n1 *int, n2 *int) {
	var n3 int = *n1
	*n1 = *n2
	*n2 = n3
}

func main() {
	var n1 int = 1
	var n2 int = 2
	swap(&n1, &n2)
	fmt.Println(n1, n2)

}
