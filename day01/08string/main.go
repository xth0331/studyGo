package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	// \本来是具有特殊含义的，需要转义
	path := "\"/Users/xieth/go/src/github.com/xth0331/studyGo/day01\""
	fmt.Printf(path)

	s := "I'm ok"
	fmt.Println(s)
	// 多行字符串
	s2 := `
		阿斯顿发
		adsf
	大法师地方
	`
	fmt.Println(s2)
	s3 := `/Users/xieth/go/src/github.com/xth0331/studyGo/day01`

	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "xiet"
	world := "adf"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	// fmt.Printf("%s%s", name,world)
	fmt.Println(ss1)
	// 分割
	ret := strings.Split(s3, "/")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "xiet"))
	fmt.Println(strings.Contains(ss, "ads"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "xiet"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "xiet"))
	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
	
}
