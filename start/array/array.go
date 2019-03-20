package main

import "fmt"

func main() {

	// array2 := [5]int{1, 2, 3, 4, 5}
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("索引:%d，值: %d\n", i, array2[i])
	// }
	//在函数间传递变量时，总是以值的方式，如果变量是个数组，那么就会整个复制，并传递给函数
	array := [5]int{1: 2, 3: 4}
	//只是复制数组，并没有修改原来的值
	fmt.Println("array before: ", array)
	modify(array)
	fmt.Println("array after: ", array)

	//传递指针，会修改数组中的值
	modify1(&array)
	fmt.Println(array)

	//slice append
	fmt.Println("---------------")
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	s2 := append(s1, 3, 5, 7)
	//两个slice合并的办法
	s3 := append(s2, s0...)
	fmt.Println(s3)
	//可以看到slice append之后的地址没有指向原来的那个
	//打印出地址需要用printf
	fmt.Printf("s0的地址: %p, s3的地址：%p", &s0, &s3)

	fmt.Println("---------------")
	//slice copy
	var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	var s = make([]int, 6)
	//copy函数返回它复制的元素的个数
	//可以使用”_“来充当匿名变量
	_ = copy(s, a[0:])
	n2 := copy(s, s[2:])
	fmt.Println("n1: ", n2)
	fmt.Println("s: ", s)
}

func modify(a [5]int) {
	a[1] = 3
	fmt.Println(a)
}

//这个地方传递的是数组的指针，不同于指针数组[5]*int
func modify1(a *[5]int) {
	a[1] = 3
	fmt.Println(*a)
}
