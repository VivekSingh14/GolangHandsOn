package main

import "fmt"

func main() {

	//declaration and initialization
	myMap := map[int]int{
		01: 101,
		02: 102,
	}

	fmt.Println(myMap)

	fmt.Println("****************")
	//assigning some key values to map
	myMap[03] = 103

	fmt.Println(myMap)

	fmt.Println("************traversing a map***********")
	//traversing a map
	for i := 0; i < len(myMap); i++ {
		fmt.Println(myMap[i+1])
	}

}
