package main

import "fmt"

func main() {
	num := 5
	temp := num
	fact := 1
	for num > 0 {
		fact = fact * num
		num = num - 1
	}

	fmt.Println("Factorial of number ", temp, " is ", fact)
}
