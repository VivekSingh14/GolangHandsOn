package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	name     string
	password string
	permissions
}

func main() {

	users := []user{
		{"Vivek", "1234", nil},
		{"Vikas", "42", permissions{"admin": true}},
		{"Vishal", "666", permissions{"write": true}},
	}

	out, err := json.Marshal(users)

	if err != nil {
		fmt.Println("err")
		return
	}

	//had to convert the out into string because by default out's type is byte.
	//and in this case we are not able to see the data because json package only encodes exported fields.
	fmt.Println(string(out))

}
