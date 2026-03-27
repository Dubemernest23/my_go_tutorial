package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Grade string
	Score int
}

// variadic function
// The ...int means "zero or more ints"
// Inside the function, scores is just a []int
// this basic variadic function

func total(scores ...int) int {
	sum := 0
	for _, s := range scores {
		sum += s
	}
	return sum
}

// usig struct
// Practical: print a report for any number of students

func printReport(title string, students ...Student) {
	fmt.Println("===", title, "===")
	for _, s := range students {
		fmt.Printf("%-15s Grade: %s  Score: %d\n", s.Name, s.Grade, s.Score)
	}
	fmt.Println("Total students:", len(students))
}

// mixed param
// Regular param first, variadic last
func gradeReport(className string, passing int, scores ...int) {
	fmt.Printf("Class: %s | Pass mark: %d\n", className, passing)
	for i, score := range scores {
		status := "FAIL"
		if score >= passing {
			status = "PASS"
		}
		fmt.Printf("  Student %d: %d — %s\n", i+1, score, status)
	}
}

// IIFE - immediately invocke funtion
// Define and call immediately — note the () at the end

func main() {

	// variadic
	printReport("Term 1 Results",
		Student{"Amara Obi", "A", 92},
		Student{"Chidi Nwosu", "B", 78},
		Student{"Fatima Bello", "A", 95},
	)
	fmt.Println(total())           // 0   — zero args is fine
	fmt.Println(total(80))         // 80
	fmt.Println(total(80, 90, 95)) // 265
	gradeReport("SS3A", 50, 92, 78, 45, 61, 38)

	// anonymous

}
