package main

import "fmt"

func main() {
	age := 10
	if age < 18 {
		fmt.Println("Kid")
	} else {
		fmt.Println("Adult")
	}
	sum := 0
	for i := 0; i < 100; i++ {
		if sum > 100 {
			break
		}
		sum += i
	}
	fmt.Println(sum)
	nums := []int{10, 20, 30, 40}
	for i, num := range nums {
		fmt.Println(i, num)
	}

	gender := 2

	switch gender {
	case 1:
		fmt.Println("female")
	case 2:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
}
