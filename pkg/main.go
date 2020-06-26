package main

import "fmt"

//注意在引用本地包的定位格式
import "pkg/mathop"

func main() {
	fmt.Println(mathop.Multi(8, 9)) // 72
}
