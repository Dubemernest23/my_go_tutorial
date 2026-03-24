package main

import (
	"fmt"
)

func main() {

	user := []struct {
		id       int
		username string
		state    string
		rating   float64
		isActive bool
	}{
		{id: 4599, username: "emeka3", state: "enugu", rating: 2.4, isActive: false},
		{id: 4543, username: "e3tun", state: "imo", rating: 1.4, isActive: true},
		{id: 4519, username: "akunne", state: "enugu", rating: 4.2, isActive: false},
		{id: 4269, username: "ekenq2", state: "abia", rating: 3.1, isActive: true},
	}
	total := 0.0

	for i, v := range user {
		total = total + v.rating
		fmt.Println("Person", i, "Rating", v.rating)
		fmt.Printf("%v, has the rating of %v \n", v.username, v.rating)
	}
	fmt.Println(total)

	todos := []string{
		"learn golang",
		"work out",
		"play games",
	}

	names := []string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	todos = append(todos, names...)

	// fmt.Println(user)
	// fmt.Println(todos)
}
