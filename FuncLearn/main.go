package main

import (
	"fmt"
	"os"
)

func main() {

	result := multiply(os.Args[1], os.Args[2])

	fmt.Println("multiplication of ", os.Args[1], " and ", os.Args[2], result)

}

func multiply(a, b string) string {
	return a + b
}
