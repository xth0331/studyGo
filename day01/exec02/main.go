package main

import (
	"unicode/utf8"
	"fmt"
)

var name string = "hello沙河小王子"

func main() {
	hansCount := utf8.RuneCountInString(name)
	fmt.Printf("name字符串中中文的数量是：%v\n", hansCount)
}
