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
}
