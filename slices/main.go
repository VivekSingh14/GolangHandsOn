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
}
