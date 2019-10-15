package main

import (
	"fmt"
	"strings"
)

func main() {
	address := "1.1.1.1"
	fmt.Println(defangIPaddr(address))
}

func defangIPaddr(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func defangIPaddr2(address string) string {
	a := "."
	b := "[.]"
	rep := ""
	for _, s := range address {
		if string(s) == a {
			rep += b
			continue
		}
		rep += string(s)
	}
	return rep
}
