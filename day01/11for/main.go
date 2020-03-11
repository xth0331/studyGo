package main

import "fmt"

func main() {
	// 基本格式
	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// }
	// 变种1
	// var i = 5
	// for; i <= 10; i++{
	// 	fmt.Println(i)
	// }
	// 变种2
	// var i = 5
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	// 无限循环
	// for {
	// 	fmt.Println("123")
	// }
	// for range 用来遍历，返回索引 值
	s := "Hello沙河"
	for i,v := range s{
		fmt.Printf("%d %c\n", i, v)
	}

}