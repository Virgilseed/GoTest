package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id int
}

func main() {
	var s1 Student = Student{Person{"mike", 'm', 18}, 1}
	fmt.Println("s1 = ", s1)
}
