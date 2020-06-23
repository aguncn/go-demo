package main

import "fmt"

func power(num int) {
	num *= num
}

func realPower(num *int) {
	*num *= (*num)
}

func main() {
	str := "Golang"
	var p *string = &str // p 是指向 str 的指针
	*p = "Hello"
	fmt.Println(str) // Hello 修改了 p，str 的值也发生了改变
	num := 100
	power(num)
	fmt.Println(num) // 100，num 没有变化

	realPower(&num)
	fmt.Println(num) // 101，指针传递，num 被修改
}
