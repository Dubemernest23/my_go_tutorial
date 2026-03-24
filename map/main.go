package main

import "fmt"

func main() {
	x := map[string]string{
		"username": "emeka",
		"password": "password",
	}

	fmt.Println(len(x))
	fmt.Println(x["username"])
}
