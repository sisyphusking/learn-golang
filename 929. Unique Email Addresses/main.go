package main

import (
	"fmt"
	"strings"
)

func main() {

	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	//这里原本是想有个set的集合的，最后返回set的长度；
	//但是golang中是没有这个数据类型，有个很奇妙的设计就是用map，把不同的元素当做key来存进去，最后统计出map的长度
	seen := make(map[string]bool)

	for _, email := range emails {
		comp := strings.Split(email, "@")
		comp[0] = strings.Replace(comp[0], ".", "", -1)
		comp[0] = strings.Split(comp[0], "+")[0]
		seen[strings.Join(comp, "")] = true
	}
	fmt.Println(len(seen))
}
