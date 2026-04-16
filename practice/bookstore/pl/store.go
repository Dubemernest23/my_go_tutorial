package pl

import "fmt"

type Book struct {
	// ID, Title, Author, Year, Available (bool)
	ID        int
	Title     string
	Author    string
	Year      int
	Available bool
}

type Library struct {
	Books  []*Book
	nextID int
}

type IlibraryInterface interface {
	AddBook(book *Book) error
	RemoveBook(id int) error
	FindByID(id int) (*Book, error)
	ListBooks() []*Book
	CheckoutFunc(id int) error
	ReturnFunc(id int) error
}

var _ IlibraryInterface = (*Library)(nil)

func NewBook(title string, author string, year int) *Book {
	return &Book{
		Title:     title,
		Author:    author,
		Year:      year,
		Available: true,
	}
}

func NewLibrary() *Library {
	return &Library{
		Books:  []*Book{},
		nextID: 1,
	}
}

// add method that is attached to the Library struct, which takes a Book as an argument and
// adds it to the Books slice.
// The method should also assign a unique ID to each book added to the library.
func (l *Library) AddBook(book *Book) error {
	book.ID = l.nextID
	l.nextID++
	l.Books = append(l.Books, book)
	return nil
}

// list method that returns a slice of all the books in the library.
//
//	This method should be attached to the Library struct and should return a
//
// slice of pointers to Book structs.
func (l *Library) ListBooks() []*Book {
	return l.Books
}

// remove method that takes a book ID as an argument and
// removes the corresponding book from the library.
func (l *Library) RemoveBook(id int) error {
	for i, b := range l.Books {
		if b.ID == id {
			l.Books = append(l.Books[:i], l.Books[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("book with ID %d not found", id)
}

// findbyid method that takes a book ID as an argument and returns the
// corresponding book from the library.
func (l *Library) FindByID(id int) (*Book, error) {
	for _, b := range l.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, fmt.Errorf("book with ID %d not found", id)
}

// checkout method that takes a book ID as an argument and marks the corresponding book
//  as unavailable.
func (l *Library) CheckoutFunc(id int) error {
	book, err := l.FindByID(id)
	if err != nil {
		return err
	}

	if book.Available == false {
		return fmt.Errorf("book with ID %d is already checked out", id)
	}

	book.Available = false
	return nil
}

// return method that takes a book ID as an argument and marks the corresponding book as available.
func (l *Library) ReturnFunc(id int) error {
	book, err := l.FindByID(id)
	if err != nil {
		return err
	}
	if book.Available == true {
		return fmt.Errorf("book with ID %d is already available", id)
	}
	book.Available = true
	return nil
}
