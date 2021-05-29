package main

import "fmt"

type AWSMachine struct {
	Name         string
	instanceType string
	flavour      string
	memory       int
	started      bool
}

func main() {

	fmt.Println("Vivek Singh")

	fmt.Println("Shahrukh Khan")

	fmt.Println("VIKAS SINGH")

	str := test(2)
	fmt.Printf("%s \n", str)

	fmt.Println("**************SLICES***********")

	var ages []int32
	for i := 0; i < 4; i++ {
		ages = append(ages, int32(i)+2)
	}
	ages = append(ages, 12)
	fmt.Println(ages)
	for _, test := range ages {
		fmt.Printf("%d ", test)
	}
	fmt.Println()
	for i := 0; i < len(ages); i++ {
		fmt.Print(ages[i], " ")
	}
	fmt.Println()

	fmt.Println("***************POINTER*****************")
	//Pointer learning

	var vivek int = 100

	P := &vivek
	fmt.Printf("%d\n%d\n%d\n", P, &P, *P)

	fmt.Println("********************Structure*********************")

	instance1 := AWSMachine{
		Name:         "webserver",
		instanceType: "EC2",
		flavour:      "Ubuntu18.04",
		memory:       40,
		started:      false,
	}
	fmt.Println(instance1.Name, instance1.flavour)

	fmt.Println("******************MAPS*****************************")
	map1 := map[int]string{
		90: "Vivek",
		91: "Singh",
		92: "Rajat",
		93: "Sharma",
		94: "Prateet",
		95: "Jain",
	}
	fmt.Println("Map-1: ", map1)
	fmt.Println("*************************")

	for per, name := range map1 {
		fmt.Printf("%d %s \t", per, name)
	}
	fmt.Println()
}

func test(one int) string {
	fmt.Printf("my rank is %d \n", one)
	return "Vivek is best"
}
