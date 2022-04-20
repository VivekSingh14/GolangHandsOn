package main

import "fmt"

func main() {

	str := "abcbdeafedz"
	arr := []rune(str)

	map1 := make(map[rune]int)

	for i := range arr {
		if _, ok := map1[arr[i]]; ok {
			temp := map1[arr[i]]
			temp += 1
			map1[arr[i]] = temp
		} else {
			map1[arr[i]] = 1
		}
	}

	for r, v := range map1 {
		if v == 2 {
			fmt.Printf("%c \t %d \n", r, v)
		}
	}

}
