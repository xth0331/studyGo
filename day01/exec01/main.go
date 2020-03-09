package main

import "fmt"

var (
	n1 int64 = 100
	n2 float64 = 3.14
	n3 bool = false
	n4 string = "Hello"
)

func main() {
	fmt.Printf("n1:value is %v, type is %T\n", n1, n1)
	fmt.Printf("n2:value is %v, type is %T\n", n2, n2)
	fmt.Printf("n3:value is %v, type is %T\n", n3, n3)
	fmt.Printf("n4:value is %v, type is %T\n", n4, n4)
}

