package main

import (
	"fmt"
)

// struct method

type Address struct {
	City  string
	State string
}

// Nested struct
type User struct {
	Base
	Username string
	Rating   float64
	IsActive bool
	Address  Address // Nested struct
}

// Struct Embedding (like inheritance)
// Go doesn't have inheritance,
// but you can embed one struct inside another to reuse its fields and methods.
type Base struct {
	ID        int
	CreatedAt string
}

type Product struct {
	Base  // Product also has ID and CreatedAt
	Name  string
	Price float64
}

// An interface defines behaviour, a struct implements it —
// just by having the matching methods
type Animal interface {
	Sound() string
	Name() string
}

type Dog struct{ Breed string }
type Cat struct{ Color string }

func (d Dog) Sound() string { return "Woof" }
func (d Dog) Name() string  { return "Dog" }

func (c Cat) Sound() string { return "Meow" }
func (c Cat) Name() string  { return "Cat" }

// this function accepts ANY type that satisfies the Animal interface
func describe(a Animal) {
	fmt.Printf("I am a %s and I go %s\n", a.Name(), a.Sound())
}

// (u User) is the receiver — like 'this' in TypeScript
func (u User) Greet() string {
	return "Hi, I am " + u.Username
}

// pointer reciever
// If you want a method to modify the struct, you need a pointer receiver *User.
// Otherwise Go passes a copy and your changes are lost.
func (u *User) Activate() {
	u.IsActive = true
}

func main() {
	u := User{Base: Base{ID: 1, CreatedAt: "2024-01-01"}, Username: "emeka", IsActive: false, Rating: 4.5, Address: Address{
		City:  "okunano",
		State: "enugu",
	}}

	u.Activate()

	describe(Dog{Breed: "Bully"}) // I am a Dog and I go Woof
	describe(Cat{Color: "Black"}) // I am a Cat and I go Meow

	fmt.Println(u.ID)
	fmt.Println(u.IsActive)      // false
	fmt.Println(u.IsActive)      // true
	fmt.Println(u.Greet())       // Hi, I am emeka
	fmt.Println(u.Username)      // emeka
	fmt.Println(u.Address.City)  // Enugu
	fmt.Println(u.Address.State) // Enugu
}
