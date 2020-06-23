package main

import "fmt"

func main() {
	//数组声明时初始化
	array := [5]int{10, 12, 13, 14, 15}
	for i := 0; i < len(array); i++ {
		array[i] += 100
	}
	fmt.Println(array)

	// [0 0 0] 长度为3容量为5的切片
	slice := make([]float32, 3, 5)
	fmt.Println(len(slice), cap(slice), slice)
	// 添加元素，切片容量可以根据需要自动扩展
	slice = append(slice, 10, 20, 30, 40)
	fmt.Println(len(slice), cap(slice), slice)
}
