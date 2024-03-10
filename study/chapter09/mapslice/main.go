package main

import "fmt"

// 演示map slice的使用
func main() {

	//声明一个切片
	var monsters []map[string]string = make([]map[string]string, 2)
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "牛魔王"
		monsters[0]["age"] = "500"
	}

	if monsters[1] == nil {
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "狐狸精"
		monsters[1]["age"] = "200"
	}

	//使用切片的append来动态增加
	//新创建一个monster信息
	newMonster := map[string]string{
		"name": "新的妖怪～",
		"age":  "300",
	}
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

}
