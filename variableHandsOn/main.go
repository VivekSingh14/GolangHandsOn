package main

import "fmt"

func main() {
	var temp int
	var temp1 float64
	var temp2 bool
	var temp3 string
	//temp = 2

	fmt.Println(temp, temp1, temp2, temp3)
	//array declaration
	var arr [4]int
	arr[0] = 5
	arr[1] = 6
	arr[2] = 7
	arr[3] = arr[1] + 8

	fmt.Println(arr)
	fmt.Println(arr[3])

	//slice declaration
	arr1 := []int{1, 2, 3, 4}

	//or

	var arr2 []int

	arr2 = append(arr2, 1)
	arr2 = append(arr2, 2)
	arr2 = append(arr2, 3)
	arr2 = append(arr2, 4)
	arr2 = append(arr2, 5)

	fmt.Println("+++++++++++++++++++++++++++++++")
	fmt.Println(arr2)

	//arr1[0] = 2
	//arr1[1] = 4
	//arr1[2] = 4
	//arr1[3] = 4
	//arr1[4] = 4

	fmt.Println("****************")
	fmt.Printf("%#v\n", arr1)
	fmt.Println(arr1)
}
