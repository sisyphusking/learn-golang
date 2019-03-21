package main

import "fmt"

//int类型的type，支持int类型的运算
type data int

func (d data) String() string {
	return fmt.Sprintf("data:%d", d)
}

func main() {
	var d data = 15
	var x interface{} = d

	//<目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言
	if n, ok := x.(fmt.Stringer); ok { // 转换为更具体的接口类型
		fmt.Println(n) // 由于n实现了String()方法，所以这里打印会输出
	}

	if d2, ok := x.(data); ok { // 转换回原始类型
		fmt.Println(d2)
	}

	e := x.(error) // 错误:main.data is not error
	fmt.Println(e)
}
