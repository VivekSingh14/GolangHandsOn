package main

import "fmt"

func main() {
	go display("Welcome")

	display("Vivek")
}

func display(str string) {
	for i := 0; i < 6; i++ {
		fmt.Println(str)
	}
}
