package main

import "fmt"

func modifyUser(users map[string]map[string]string, name string) {

	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称" + name
	}
}

func main() {

	user := make(map[string]map[string]string)
	user["smith"] = make(map[string]string)
	user["smith"]["pwd"] = "123424"
	user["smith"]["nickname"] = "fefe"
	modifyUser(user, "tom")
	modifyUser(user, "mary")
	modifyUser(user, "smith")
	fmt.Println(user)
}
