package main

import (
	"fmt"
	"time"
)

func main() {
	go display("Welcome")

	display("Vivek")
}

func display(str string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}
