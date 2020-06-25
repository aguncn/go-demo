package main

import "fmt"

type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

func (w *Worker) getName() string {
	return w.name
}

func main() {
	var p1 Person = &Student{
		name: "Jerry",
		age:  25,
	}
	var p2 Person = &Worker{
		name:   "Jack",
		gender: "male",
	}

	fmt.Println(p1.getName())
	fmt.Println(p2.getName())
}
