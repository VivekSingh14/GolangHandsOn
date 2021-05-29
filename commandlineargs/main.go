package main

import "fmt"
import "os"

func main() {

	var (
		one   int
		two   string
		three float64
	)

	one = 1
	two = "Vivek"
	three = 2.3

	fmt.Printf("%T\n", one)
	fmt.Printf("%q\n", two)
	fmt.Printf("%#v\n", two)
	fmt.Printf("%#v\n", three)

	fmt.Println(os.Args[0])
	//fmt.Println(os.Args[1])
	fmt.Println(len(os.Args))

}
