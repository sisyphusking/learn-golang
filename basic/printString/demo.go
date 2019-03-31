package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str = "hello世界"
	//golang中string底层是通过byte数组实现的。中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。
	//所以这里是13
	fmt.Println("len(str): ", len(str))
	byt := []byte(str)
	fmt.Println("byt: ", byt)

	//golang中的unicode/utf8包提供了用utf-8获取长度的方法
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))

	//通过rune类型处理unicode字符，rune等同于int32,常用来处理unicode或utf-8字符
	//[]rune 是go内建的函数，会将字符串按utf8编码转换为{h,e,l,l,o,世,界}对应的数字{104,101,108,108,111,19990,30028}
	sli := []rune(str)
	fmt.Println("sli: ", sli)
	fmt.Println("rune:", len(sli))

	//使用string()返回对应的字符
	fmt.Println("byte返回字符串：", string(byt))
	fmt.Println("rune返回字符串：", string(sli))
}
