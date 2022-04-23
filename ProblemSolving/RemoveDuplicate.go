package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 3, 4, 5, 5, 6, 8, 8}
	j := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[j] {
			j += 1
			arr[j] = arr[i]
		}
	}

	for i := 0; i < j+1; i++ {
		fmt.Printf("%d \t", arr[i])
	}

	fmt.Println()
}
