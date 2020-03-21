package main

import "fmt"

var a1 []int
var sum int

func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	// 把数组中的每个元素遍历出来，累加到一个sum变量就可以了
	sum := 0
	for _, v := range a1 {
		sum += v
	}
	fmt.Println(sum)
	// 找出和为8地两个元素地下标分别为
	// 定义两个for循环，外层的从第一个开始遍历
	// 内层的for循环从外层后面地那个开始找
	// 两个数的和为8
	for i := 0; i <= len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}

	}
}
