package main

import "fmt"

type User struct {
	ID       int
	Username string
	State    string
	Rating   float64
	IsActive bool
}

var seedUsers = []User{
	{ID: 50, Username: "mikke", State: "enugu", Rating: 4.3, IsActive: false},
	{ID: 53, Username: "chois", State: "enugu", Rating: 2.3, IsActive: false},
	{ID: 51, Username: "rene", State: "imo", Rating: 3.3, IsActive: true},
}

func getUsers() []User {
	return seedUsers
}

func printUsers(users []User) {
	for _, u := range users {
		fmt.Printf("ID: %d | Username: %-10s | State: %-6s | Rating: %.1f | Active: %v\n",
			u.ID, u.Username, u.State, u.Rating, u.IsActive)
	}
}

func main() {
	users := getUsers()
	printUsers(users)
}
