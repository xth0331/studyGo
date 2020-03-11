package main

import (
	"fmt"
)

var name = []rune("hello沙河小王子")
var count int64

func main() {
	for _, asciiValue := range name {
		if asciiValue > 127 {
			count++
		}
	}
	fmt.Printf("%v", count)
	// hansCount := utf8.RuneCount(name)
	// fmt.Printf("name字符串中中文的数量是：%v\n", hansCount)
}
