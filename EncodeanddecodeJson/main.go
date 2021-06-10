package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {

	//if we want to print name as username on result
	Name string `json:"username"`
	//if you want to hide the sensitive data such as password put json:"-", put opposite to it.
	Password string `json:"-"`

	//added omitempty because if some permissions or params having null value in any variable it will skip that variables param or field.
	Permissions permissions `json:"params,omitempty"`
}

func main() {

	users := []user{
		{"Vivek", "1234", nil},
		{"Vikas", "42", permissions{"admin": true}},
		{"Vishal", "666", permissions{"write": true}},
	}

	//out, err := json.Marshal(users)

	out, err := json.MarshalIndent(users, "", "\t")

	if err != nil {
		fmt.Println("err")
		return
	}

	//had to convert the out into string because by default out's type is byte.
	//and in this case we are not able to see the data because json package only encodes exported fields.

	//so if we want to print the json result we need to set struct variable names starting from upper case.
	fmt.Println(string(out))

}
