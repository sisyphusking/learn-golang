package main

import "fmt"

func main() {
	name := make(map[string]int)
	name["xavier"] = 28

	friends := map[string]int{"xingwei": 29}

	fmt.Print(friends)
	age := name["xavier"]
	fmt.Println()
	fmt.Print(age)

	for key, value := range name {
		fmt.Println()
		fmt.Println(key, value)
	}
	// 函数间传递map是会修改原来的map里的值的
	update(name)
	fmt.Println(name["xavier"])

}

func update(a map[string]int) {
	a["xavier"] = 10
}
