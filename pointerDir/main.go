package main

import "fmt"

func main() {

	val1 := 23

	//val2 := 26

	//var p1 *int

	p1 := &val1

	p2 := &val1

	res := &p1 != &p2

	fmt.Println(res)

	res1 := p1 == p2

	fmt.Println(res1)
}
