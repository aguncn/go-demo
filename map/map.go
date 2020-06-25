package main

import "fmt"

func main() {
	// 仅声明
	map1 := make(map[string]int)
	// 声明时初始化
	map2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}
	// 赋值/修改
	map1["Tom"] = 18

	fmt.Println(map1)
	fmt.Println(map2)
}
