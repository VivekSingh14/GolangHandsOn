package main

import "fmt"

func main() {

	//declaration and initialization
	myMap := map[int]int{
		01: 101,
		02: 102,
	}

	fmt.Println(myMap)

	fmt.Println("\n****************")
	//assigning some key values to map
	myMap[03] = 103

	fmt.Println(myMap)

	fmt.Println("\n************traversing a map***********")
	//traversing a map
	for i := 0; i < len(myMap); i++ {
		fmt.Println(myMap[i+1])
	}

	fmt.Println("\n***********traversing a map using for each or range loop********")

	for id, roll := range myMap {
		fmt.Println(id, " : ", roll)
	}

	fmt.Println("\n*******check the existance of key in map*******")

	roll, ok := myMap[01]

	fmt.Println("key is present: ", ok)
	fmt.Println("and Value for that key is: ", roll, " \n ")

	//03 deleted here.
	fmt.Println("*********03 deleted here and checking whether it is deleted or not !**********")
	delete(myMap, 03)

	_, ok1 := myMap[03]

	fmt.Println("is", 04, "available in the map ? ", ok1)

	//modifying the older map.
	newMap := myMap

	newMap[04] = 104
	newMap[05] = 105

	fmt.Println("\n**********after adding the key values to it.*********** ")
	fmt.Println("\n", myMap)

}
