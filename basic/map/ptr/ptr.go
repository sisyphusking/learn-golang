package main

import "fmt"

type person struct {
	age  int
	name string
}

func main() {

	// 验证map类型是否会修改原值
	age := map[string]int{"xavier": 28}
	fmt.Println(age)
	modify(age)
	fmt.Println(age)

	//验证struct类型是否会修改原值
	//也可以这样赋值，但是顺序一定要和struct对应
	// myself := person(28,"xavier")
	//这种就不需要一一对应了
	myself := person{name: "xavier", age: 28}
	fmt.Println(myself)
	//传递结构体不会修改原来的值
	modify1(myself)
	fmt.Println(myself)

	//传递的是结构体的指针
	myself1 := person{name: "xavier", age: 28}
	fmt.Println(myself1)
	modify2(&myself1)
	fmt.Println(myself1)

}

//map类型传递的是引用类型，会修改原来的值
func modify(m map[string]int) {
	m["xavier"] = 18
}

//验证结构体，结构体传递的是其本身以及里面的值的拷贝，不会修改原来的值
func modify1(p person) {
	p.age = 18
}

// 传递结构体的指针，会修改原值
func modify2(p *person) {
	p.age = 18
}
