package main

// this import brings in the packages you need to use in your code.
// In this case, we are importing the "fmt" package,
//  which provides functions for formatted I/O, such as printing to the console.

import (
	"fmt"
	"math"
	"strings" // this brings in the string packahge and when its called, it beings all the methods associated
)

func main() {
	username, isActive := "sommy", true // short variable declaration with multiple variables,
	// type is inferred and cannot be used outside of a function,
	// and cannot be redeclared in the same scope. it can only be used for var and not for const

	fmt.Printf("Hello, %v! Welcome to Go programming!\n", strings.ToUpper(username))
	fmt.Printf("User %s is active: %t\n", strings.ToUpper(username), isActive)

	var radius float64 = 5.0
	area := calculateCircleArea(radius)
	fmt.Printf("The area of a circle with radius %.2f is %.2f\n", radius, area)

	num := math.Sqrt(45)
	fmt.Printf("The square root of 45 is %.2f\n", num)

	avgRating := calcAvgRating()
	fmt.Printf("The average rating is %.2f\n", avgRating)

	avgLikes := calcAverageLikes()
	fmt.Printf("The average likes is %d\n", avgLikes)
}

func calculateCircleArea(radius float64) float64 {
	math.RoundToEven(54.5) // this is a function from the math package that rounds a float64 to the nearest even integer, and returns the result as a float64.
	return math.Pi * radius * radius
}

func calcAvgRating() float64 {
	rating1 := 4.5
	rating2 := 3.8
	rating3 := 4.2
	avg := (rating1 + rating2 + rating3) / 3
	return math.Round(avg*100) / 100 // this rounds the average to 2 decimal places
}

func calcAverageLikes() int {
	like1 := 30
	like2 := 45
	like3 := 25
	avg := (like1 + like2 + like3) / 3
	return avg
}
