package main

import "fmt"

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1 // 赋值
	var a3 = make([]int, 3, 3)
	copy(a3, a1) // copy
	fmt.Println(a1, a2, a3)

	// 从切片中删除元素
	//
}
