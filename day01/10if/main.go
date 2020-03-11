package main

import "fmt"

// if 条件判断

func main() {
	// age := 19

	// if age > 18 { // 如果age > 18 就执行{}中的代码
	// 	fmt.Println("开业啦")
	// } else { // 否则就执行这个{}中的代码
	// 	fmt.Println("作业")
	// }
	// if age > 35 {
	// 	fmt.Println("中年")
	// } else if age > 18 {
	// 	fmt.Println("青年")
	// } else {
	// 	fmt.Println("好好学习")
	// }
	if age := 19; age > 18 {
		fmt.Println("开业啦")
	} else {
		fmt.Println("写作业")
	}

}
