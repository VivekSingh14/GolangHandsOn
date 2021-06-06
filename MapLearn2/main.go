package main

import (
	"fmt"
	"os"
)

func main() {

	query := os.Args[1]
	//query := "good"
	//fmt.Println(query)

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	turkish := make(map[string]string, len(dict))

	for k, v := range dict {
		turkish[v] = k
		//fmt.Println(k, " ", v)
	}

	//fmt.Println(dict, "\n", turkish)

	if value, ok := dict[query]; ok {
		fmt.Println(query, " means ", value)
		return
	}

	if value, ok := turkish[query]; ok {
		fmt.Println(query, " means ", value)
		return
	}
}
