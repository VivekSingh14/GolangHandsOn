package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(time.Now())

	const max = 5
	for found := 0; found < max; found++ {
		//n := rand.Intn(max) + 1
		//fmt.Println(n, " ")
	}
}
