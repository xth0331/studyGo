package main

import (
	"fmt"
)

// 单行注释
/*
多行
注释
*/

// 函数外的语句，必须以关键字开头
var name = "谢"
var age int

const (
	Num = 100
)

// main函数是入口函数
// 它没有参数也没有返回值
func main() {
	// 函数内定义的变量必须使用
	var isOK = true
	fmt.Println(isOK)

	// if
	var age = 19
	if age > 18 {
		fmt.Println("成年")
	} else if age > 7 {
		fmt.Println("小学")
	} else {
		fmt.Println("快乐")
	}

	// for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for2
	var i = 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	// for3
	var j = 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	// for4
	// for {
	// 	fmt.Println("无限循环")
	// }

	// for5
	s := "撒地方将阿斯顿发"
	fmt.Println(len(s))
	for i,v := range s {
		// fmt.Println(i,v)
		fmt.Printf("%d %c\n", i,v)
	}

	// 哑元变量
	for _, v := range s {
		// fmt.Println(v)
		fmt.Printf("%c\n", v)
	}
}
