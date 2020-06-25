package main

import (
	"fmt"
)

func rect(length, width float64) (float64, float64) {
	var a = length * width
	var p = (length + width) * 2
	return a, p
}

func main() {
	area, perimeter := rect(100, 50)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)
}
