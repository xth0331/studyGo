package main

import "fmt"

// map

func main() {
	var m1 map[string]int
	fmt.Println(m1 == nil)        // 还没有初始化 （没有在内存中开辟空间）
	m1 = make(map[string]int, 10) // 要估算好该map容量，避免在程序运行期间再动态扩容
	m1["实施"] = 18
	m1["jisdf"] = 35
	fmt.Println(m1)
	fmt.Println(m1["实施"])
	value, ok := m1["士大夫"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(value)
	}
	// map 遍历
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	// 仅遍历 key
	for k := range m1 {
		fmt.Println(k)
	}
	// 仅遍历 value
	for _, v := range m1 {
		fmt.Println(v)
	}
	// 删除
	delete(m1, "jisdf")
	fmt.Println(m1)
}
