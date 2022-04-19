package main

import "fmt"

func main() {
	var arr [3]int

	for i := 0; i < 3; i++ {
		arr[i] = i
	}

	for i := 0; i < 3; i++ {
		fmt.Println(arr[i])
	}
	fmt.Print("length of an array: ", len(arr), "\n")
	//fmt.Println(len(arr))

	fmt.Println("traversing the array using range inbuilt variable")
	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println("traversing Hello for commit array using for each ")
	//took boolean array of size arr
	var published [len(arr)]bool

	published[0] = true
	published[len(arr)-1] = true
	for i, ok := range published {
		if ok {
			fmt.Println(arr[i])
		}

	}

	numarr := [3]int{1, 2, 3}

	fmt.Println(numarr)

}
