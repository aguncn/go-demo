package main

import "fmt"

// 定义只有一个方法的接口
type Person interface {
	getName() string
}

type Student struct {
	name string
	age  int
}

//Student实现了接口
func (stu *Student) getName() string {
	return stu.name
}

type Worker struct {
	name   string
	gender string
}

//Worker实现了接口
func (w *Worker) getName() string {
	return w.name
}

func main() {
	//可用实现了接口的类型实例化接口
	var p1 Person = &Student{
		name: "Jerry",
		age:  25,
	}
	var p2 Person = &Worker{
		name:   "Jack",
		gender: "male",
	}
	//不同的接口实例可调用相同的接口成员方法，但输出不同
	fmt.Println(p1.getName())
	fmt.Println(p2.getName())
}
