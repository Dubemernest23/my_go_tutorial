package main

import "fmt"

func main() {
	var isActive bool = false
	var username string = "sommy"
	var age int = 25
	var id uint = 12345
	additionalInferredVar := "jjhh"
	var score float64 = 95.5
	const pi float64 = 3.14159
	x := 5 // short variable declaration, type is inferred and cannot be used outside of a function,
	//  and cannot be redeclared in the same scope. it can only be used for var and not for const
	fmt.Print(additionalInferredVar)
	fmt.Printf("The user with username %s is active: %t\n", username, isActive) // %s for string, %t for boolean
	fmt.Printf("The user is %d years old\n", age)                               // %d for integers, %f for floating-point numbers, and %v for any value in a default format.
	fmt.Printf("The user's ID is %d\n", id)                                     // %d for unsigned integers as well
	fmt.Printf("The user's score is %.1f\n", score)                             // %.1f for floating-point numbers with one decimal place
	fmt.Printf("The value of pi is %.5f\n", pi)                                 //	 %.5f for floating-point numbers with five decimal places
	fmt.Printf("The value of x is %d\n", x)                                     // %d for integers, and since x is inferred to be an int, we use %d to print it
	fmt.Printf("What about the score? %v i believe its a high score!\n", score) //	%v for any value in a default format, and since complex numbers don't have a specific format specifier, we can use %v to print them.
	checkUserStatus()
}

func checkUserStatus() bool {
	loggedIn := true
	hasSubscription := false
	isAdmin := false
	canOpenDashboard := loggedIn && hasSubscription // this is a boolean expression that evaluates to false,
	// because loggedIn is true but hasSubscription is false.
	//  The && operator returns true only if both operands are true.

	canDeleteContent := isAdmin || canOpenDashboard // this is a boolean expression that evaluates to false,
	// because isAdmin is false and canOpenDashboard is false.
	//  The || operator returns true if at least one operand is true.

	return canDeleteContent
}

// types
// bool
// string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64
// rune
// byte
// float32, float64
// complex64, complex128
