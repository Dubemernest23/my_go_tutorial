package main

import "fmt"

type students struct {
	name        string
	grade       int
	score       int
	classRating float64
	pass        bool
}

func main() {
	x := map[string]string{
		"username": "emeka",
		"password": "password",
	}

	y := map[string]students{}

	fmt.Println(len(x))
	fmt.Println(x["username"])
	fmt.Printf("map %v \n", y)
}
