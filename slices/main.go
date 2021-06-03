package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2}

	//append
	arr = append(arr, 7)

	arr = append(arr, 3, 21)
	//arr = append(arr, 21)

	fmt.Println(arr)

	fmt.Println("after sorting")

	sort.Ints(arr)

	fmt.Println(arr)

	//after slicing the slice
	arr2 := arr[0:2]

	fmt.Println(arr2)

	arr2 = arr[:4]

	fmt.Println(arr2)

	arr2 = arr[1:]
	fmt.Println(arr2)

	str := []string{"Vivek", "Singh", "Test", "Complete"}

	fmt.Println(str)
}
