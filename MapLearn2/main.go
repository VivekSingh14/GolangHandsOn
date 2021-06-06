package main

import (
	"fmt"
	"os"
)

func main() {

	query := os.Args[0]

	fmt.Println(query)

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mukemmel",
	}

	turkish := make(map[string]string, len(dict))

	for k, v := range dict {
		turkish[v] = k
	}

	fmt.Println(dict, "\n", turkish)
}
