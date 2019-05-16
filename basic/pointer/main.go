package main

import "fmt"

/*
值传递和指针传递的区别：
参考文章：https://leokongwq.github.io/2017/01/22/golang-param-pass-value-or-point.html
对于Go语言，严格意义上来讲，只有一种传递，也就是按值传递(by value)。当一个变量当作参数传递的时候，会创建一个变量的副本，然后传递给函数或者方法，你可以看到这个副本的地址和变量的地址是不一样的。

当变量当做指针被传递的时候，一个新的指针被创建，它指向变量指向的同样的内存地址，所以你可以将这个指针看成原始变量指针的副本。当这样理解的时候，我们就可以理解成Go总是创建一个副本按值转递，只不过这个副本有时候是变量的副本，有时候是变量指针的副本。
*/

type Bird struct {
	Age  int
	Name string
}

func passV(b Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p\n", b, &b)
}

func passP(b *Bird) {
	b.Age++
	b.Name = "Great" + b.Name
	fmt.Printf("传入修改后的Bird:\t %+v, \t内存地址：%p, 指针的内存地址: %p\n", *b, b, &b)
}
func main() {
	fmt.Println("pass value.....")
	parrot := Bird{Age: 1, Name: "Blue"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t内存地址：%p\n", parrot, &parrot)
	passV(parrot)
	fmt.Printf("调用后原始的Bird:\t %+v, \t\t内存地址：%p\n", parrot, &parrot)

	fmt.Println("...........")
	fmt.Println("pass pointer.....")
	parrot1 := &Bird{Age: 2, Name: "yellow"}
	fmt.Printf("原始的Bird:\t\t %+v, \t\t内存地址：%p, 指针的内存地址: %p\n", *parrot1, parrot1, &parrot1)
	passP(parrot1)
	fmt.Printf("调用后原始的Bird:\t %+v, \t内存地址：%p, 指针的内存地址: %p\n", *parrot1, parrot1, &parrot1)

	// 可以看到在函数passP中，参数b是一个指向Bird的指针，传递参数给它的时候会创建指针的副本(0xc00000c038)，
	// 只不过指针0xc00000c038和0xc00000c030都指向内存地址0xc00000a100。 函数内对*T的改变显然会影响原始的对象，因为它是对同一个对象的操作。
}

/*
pass value.....
原始的Bird:		 {Age:1 Name:Blue}, 		内存地址：0xc00000a060
传入修改后的Bird:	 {Age:2 Name:GreatBlue}, 	内存地址：0xc00000a0a0
调用后原始的Bird:	 {Age:1 Name:Blue}, 		内存地址：0xc00000a060
...........
pass pointer.....
原始的Bird:		 {Age:2 Name:yellow}, 		内存地址：0xc00000a100, 指针的内存地址: 0xc00000c030
传入修改后的Bird:	 {Age:3 Name:Greatyellow}, 	内存地址：0xc00000a100, 指针的内存地址: 0xc00000c038
调用后原始的Bird:	 {Age:3 Name:Greatyellow}, 	内存地址：0xc00000a100, 指针的内存地址: 0xc00000c030
*/
