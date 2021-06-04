package main

import "fmt"

type Student struct {
	name  string
	class int
	age   int
}

func main() {

	var one Student

	one.name = "Vivek"
	one.class = 12
	one.age = 17

	fmt.Println(one)

	two := one.nameChange()

	fmt.Println(two)

	three := ageChange(two)

	fmt.Println(three)

	four := Student{name: "Vishal", class: 12, age: 19}

	fmt.Println(four)

}

func (s Student) nameChange() Student {
	s.name = "Singh"

	return s
}

func ageChange(s Student) Student {
	s.age = 18
	return s
}
