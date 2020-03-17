package main

import "fmt"

// 流程控制之跳出for循环

func main() {
	// 当i = 5 时，就跳出for循环
	for i:=0; i<10; i++{
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")
	// 当i = 5 ,跳过此次循环（不执行for循环内部的语句）,继续下次循环
	for i := 0 ; i<=10;i++ {
		if i == 5 {
			continue // 继续下一次循环
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
