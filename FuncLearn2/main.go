package main

import (
	"fmt"
	"os"
)

func main() {
	str1 := os.Args[1]

	str2 := os.Args[2]

	result := additionString(str1, str2)

	fmt.Println("Integers after addition ", result)

	resultInt := additionInt(435, 430)

	fmt.Println("Integers after addition ", resultInt)
}
