package main

import "fmt"

func main() {

<<<<<<< HEAD
	str := "abcbdeafedz"
=======
	str := "abcbdeafedi"
>>>>>>> 8035e4adb83db7910efb5177f083012ce261daee
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
