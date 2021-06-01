package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr []int

	//append
	arr = append(arr, 7)
	arr = append(arr, 3)
	arr = append(arr, 21)

	fmt.Println(arr)

	fmt.Println("after sorting")

	sort.Ints(arr)

	fmt.Println(arr)
}
