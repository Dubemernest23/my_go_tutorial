package main

import (
	"bookstore_api/pl"
	"fmt"
)

func main() {
	fmt.Println("=== Bookstore Library System Demo ===")

	// Create a new library
	// lb := pl.NewLibrary()
	// var lb pl.IlibraryInterface = pl.NewLibrary()
	// var lb  pl.IlibraryInterface = pl.NewLibrary()/
	var lb pl.IlibraryInterface = pl.NewLibrary()

	// Create books using the constructor
	book1 := pl.NewBook("The Great Gatsby", "F. Scott Fitzgerald", 1925)
	book2 := pl.NewBook("To Kill a Mockingbird", "Harper Lee", 1960)
	book3 := pl.NewBook("The Heist", "John Doe", 2020)
	book4 := pl.NewBook("The Trial", "Joae Dote", 2025)
	// Add books to the library
	fmt.Println("Adding books to the library...")
	lb.AddBook(book1)
	lb.AddBook(book2)
	lb.AddBook(book3)
	lb.AddBook(book4)

	fmt.Println("Books added successfully!")

	// List all books
	fmt.Println("Current books in library:")
	for _, b := range lb.ListBooks() {
		fmt.Printf("ID:%d | %s by %s (%d) | Available: %v\n",
			b.ID, b.Title, b.Author, b.Year, b.Available)
	}

	// Checkout a book
	fmt.Println("\n--- Checking out book ID 1 ---")
	err := lb.CheckoutFunc(1)
	if err != nil {
		fmt.Println("Checkout Error:", err)
	} else {
		fmt.Println("Book ID 1 checked out successfully.")
	}

	// Try to checkout the same book again (should fail)
	fmt.Println("\n--- Trying to checkout book ID 1 again ---")
	err = lb.CheckoutFunc(1)
	if err != nil {
		fmt.Println("Checkout Error:", err)
	}

	// Return the book
	fmt.Println("\n--- Returning book ID 1 ---")
	err = lb.ReturnFunc(1)
	if err != nil {
		fmt.Println("Return Error:", err)
	} else {
		fmt.Println("Book ID 1 returned successfully.")
	}

	// Remove a book
	fmt.Println("\n--- Removing book ID 3 ---")
	err = lb.RemoveBook(3)
	if err != nil {
		fmt.Println("Remove Error:", err)
	} else {
		fmt.Println("Book ID 3 removed successfully.")
	}

	// Final list of books
	fmt.Println("\nFinal books in library:")
	for _, b := range lb.ListBooks() {
		fmt.Printf("ID:%d | %s by %s (%d) | Available: %v\n",
			b.ID, b.Title, b.Author, b.Year, b.Available)
	}

	fmt.Println("\n=== Demo Completed ===")
}
