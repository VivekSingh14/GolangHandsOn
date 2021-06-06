package main

import "fmt"

type person struct {
	name string
	age  int
	city string
}

func main() {

	var obj1 person
	var obj2 person

	fmt.Println(obj1, "  ", obj2)

	obj1.name = "Vivek"
	obj1.city = "Jaipur"
	obj1.age = 27

	obj2.name = "Vivek2"
	obj2.city = "Pune"
	obj2.age = 25

	fmt.Println(obj1, " ", obj2, " \n ")

	map1 := map[int]person{}
	map1[1] = obj1
	map1[2] = obj2

	//whole map print
	fmt.Println(map1, " \n ")

	//map printed based on condition
	for k, v := range map1 {
		if v.name == "Vivek" {
			fmt.Println(k, " ", v)

		}
	}

}
