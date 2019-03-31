package main

import (
	"fmt"
	"time"
)

var x int = 1

func main() {

	func(a, b int) {
		fmt.Println(a + b)
	}(1, 2) //花括号后面直接加括号，表示直接调用，并且可以传入参数

	//例子2
	f := outer(10)
	fmt.Println("result is :", f(100))

	//例子3
	s := []string{"a", "b", "c"}
	for _, v := range s {
		go func() {
			fmt.Println(v)
		}()
	}
	time.Sleep(time.Second * 1)

	//例子4
	for _, v := range s {
		go func(v string) {
			fmt.Println(v)
		}(v) //每次将变量v的拷贝传进函数，这里有个坑，不能使用(),要用(v),要不然打印输出的全部都是c，在没有将变量v的拷贝值传进匿名函数之前，只能获取最后一次循环的值
	} // 这里才是循环结束的位置
	time.Sleep(time.Second * 1)

	//例子5
	for _, f := range test() {
		f()
	}

	//例子6
	//利用闭包来修改全局变量
	y := func() int {
		x++
		return x
	}()
	fmt.Println("main:", x, y)

	//例子7
	w, z := 1, 2

	//defer 调用会在当前函数执行结束前才被执行，这些调用被称为延迟调用
	//可以看到在执行到这里的时候，w的值是1，函数在执行完下面的流程后才执行这个，所以最后打印出w的值是1，不会是101，而z会取最后保存的那个值，所以是102
	defer func(a int) {
		fmt.Printf("w:%d, z:%d\n", a, z)
	}(w)
	w += 100
	z += 100
	fmt.Println(w, z)
}

//先实例化outer，会返回一个int类型的匿名函数
//然后传入参数，这时候就是相当于给y赋值
func outer(x int) func(int) int {
	return func(y int) int { //匿名函数
		return x + y
	}
}

func test() []func() {
	var s []func()

	for i := 0; i < 3; i++ {
		//每次 append 操作仅将匿名函数放入到列表中，但并未执行，并且引用的变量都是 i，随着 i 的改变，匿名函数中的 i 也在改变
		//所以当执行这些函数时，他们读取的都是环境变量 i 最后一次的值。解决的方法就是每次复制变量 i 然后传到匿名函数中，让闭包的环境变量不相同。
		x := i //复制变量
		s = append(s, func() {
			fmt.Println(&x, x)
		})
	}
	return s
}
