package main

import "fmt"

func main() {
	x, y := 1, 2
	swap(&x, &y)
	fmt.Println(x, y)

	w, z := 3, 4
	fmt.Println("before---&w, &z: ", &w, &z)
	swap1(&w, &z) // 这个地方传递的是变量的地址，是值拷贝，对w z没有影响
	fmt.Println("after---&w, &z: ", &w, &z)
	fmt.Println(w, z)
}

func swap(a, b *int) {
	t := *a //将a指针取值，赋值给t
	*a = *b //将b指针取值，赋给a变量指向的变量
	*b = t
}

func swap1(a, b *int) {
	b, a = a, b //将a,b进行交换
}
