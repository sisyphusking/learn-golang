package main

import (
	"fmt"
)

// 定义接口
type Handler interface {
	Do(k, v interface{})
}

//function type
type HandlerFunc func(k, v interface{})

//实现了Handler接口
//Do方法的实现是调用HandlerFunc本身，因为HandlerFunc类型的变量就是一个方法。
//f(k,v)就是调用本身
func (f HandlerFunc) Do(k, v interface{}) {
	f(k, v)
}

//遍历map
func Each(m map[interface{}]interface{}, h Handler) {
	if m != nil && len(m) > 0 {
		for k, v := range m {
			h.Do(k, v)
		}
	}
}

//传入f函数，组装成HandlerFunc类型， 每一个元素都执行function type中的方法
func EachFunc(m map[interface{}]interface{}, f func(k, v interface{})) {
	Each(m, HandlerFunc(f))
}

//HandlerFunc中的方法
func selfInfo(k, v interface{}) {
	fmt.Printf("大家好,我叫%s,今年%d岁\n", k, v)
}

func main() {
	persons := make(map[interface{}]interface{})
	persons["张三"] = 20
	persons["李四"] = 23
	persons["王五"] = 26

	EachFunc(persons, selfInfo)

}
