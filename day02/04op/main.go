package main

import "fmt"

// 运算符

func main() {
	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ // 单独对语句，不能放在=右边赋值 a = a + 1
	b-- // 单独的语句，不能放在=右边赋值 b = b - 1
	fmt.Println(a)

	// 关系运算符
	fmt.Println(a == b) // Go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b) // 不等于
	fmt.Println(a >= b) // 大于等于
	fmt.Println(a <= b) // 小于等于
	fmt.Println(a > b) // 小于
	
	// 逻辑运算符
	// 如果年龄小于18， 并且年龄小于60岁
	age :=22
	if age > 18 && age < 60 {
		fmt.Println("上班")
	} else{
		fmt.Println("不用上班")
	}

	// 如果年龄小于18 或者大于60
	if age < 18 || age > 60 {
		fmt.Println("不用上班")
	} else {
		fmt.Println("上班")
	}
}
