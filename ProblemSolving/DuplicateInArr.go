package main

import "fmt"

func main() {

	arr := []int{8, 3, 2, 4, 2, 6, 3, 8}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				fmt.Println(arr[i])
			}
		}
	}
}
