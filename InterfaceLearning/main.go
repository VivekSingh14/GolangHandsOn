package main

import "fmt"

type testInterface interface {
	fun1() int
	fun2(string) int
}

func fun1() int {
	return 14
}

func fun2(string) int {
	return 15
}

func main() {

	fmt.Println(fun1())

	fmt.Println(fun2("vivek"))
}
