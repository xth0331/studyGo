package main

import "fmt"

func main() {
	for i:= 1; i <= 9; i++ {
		for j := 1; j<= i; j++{
			fmt.Printf("%v * %v = %v\t", j,i,j*i)
			// if i == j {
			// 	fmt.Println()
			// }
		}
		fmt.Println()
	}
}