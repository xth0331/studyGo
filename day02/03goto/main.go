package main

import "fmt"

// goto

// func main() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				break
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 	}
// }
func main() {
	// 跳出多层for循环
	// var flag = false
	// for i := 0; i < 10; i++ {
	// 	for j := 'A'; j < 'Z'; j++ {
	// 		if j == 'C' {
	// 			flag = true
	// 			break // 跳出内层的for循环
	// 		}
	// 		fmt.Printf("%v-%c\n", i, j)
	// 	}
	// 	if flag {
	// 		break // 跳出for循环（外层）
	// 	}
	// }
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX // 跳到我指定的标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
XX: // 标签
	fmt.Println("over")
}
