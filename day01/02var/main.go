package main

import "fmt"

// Go语言中推荐使用小驼峰式命名
// var studnet_name string

var studnetName string

// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明

var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "谢"
	age = 16
	isOk = true
	// var heiheihei string
	// Go语言中非全局变量声明后必须使用,不使用无法编译
	fmt.Print(isOk) // 在终端输入要打印的内容
	fmt.Println()
	fmt.Printf("name:%s\n", name) // %s: 占位符 使用name这个变量的值
	fmt.Println(age)

	//声明变量同时赋值
	var s1 string = "xth"
	fmt.Println(s1)
	// 类型推导(根据值判断变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)

	// 简短变量声明,只能在函数里面用
	s3 := "嘿嘿嘿"
	fmt.Println(s3)
	// s1 := "10" 同一作用于({})中不能重复声明同名的变量
	// 匿名变量是一个特殊的变量 : _
}
