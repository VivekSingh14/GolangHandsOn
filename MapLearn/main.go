package main

import "fmt"

func main() {

	myMap := map[int]int{
		01: 101,
		02: 102,
	}

	fmt.Println(myMap)

	fmt.Println("****************")

	myMap[03] = 103

	fmt.Println(myMap)

}
