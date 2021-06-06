package main

import "fmt"

type song struct {
	title, artist string
}

type playlist struct {
	genre string
	songs []song
}

func main() {

	songs := []song{
		{title: "wonderwal", artist: "oasis"},
		{"super sonic", "oasis"},
	}

	rock := playlist{genre: "rock", songs: songs}

	fmt.Println(rock)
}
