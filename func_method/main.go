package main

import "fmt"

type Staff struct {
	name string
	age  int
}

func (s *Staff) hello(name string) string {
	return fmt.Sprintf("hello %s, I am %s, I am %d years old.", name, s.name, s.age)
}

func main() {
	staff := &Staff{
		name: "Tom",
		age:  25,
	}
	msg := staff.hello("Jerry")
	fmt.Println(msg)
}
