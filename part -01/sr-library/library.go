//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string // Title of the book
type Name string  // Name of the member

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // Total number of books
	lended int // Number of books that have been checked out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "not returned"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printMembersAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
		println("------------------")
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/ total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	// Make sure the book is part of the library
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found")
		return false
	}
	// Make sure we have enough to lend
	if book.lended == book.total {
		fmt.Println("Book not available")
		return false
	}

	// Check out the book
	book.lended++
	library.books[title] = book
	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	// Make sure the book is part of this library
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not found")
		return false
	}
	// Make sure the member checked out the book
	audit, found := member.books[title]
	if !found {
		fmt.Println("Book not checked out")
		return false
	}

	// Return the book
	book.lended--
	library.books[title] = book
	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {

	library := Library{
		members: make(map[Name]Member),
		books:   make(map[Title]BookEntry),
	}
	//  - Add at least 4 books...
	library.books["Webapps in Go"] = BookEntry{
		total:  4,
		lended: 0,
	}

	library.books["Learning Go"] = BookEntry{
		total:  3,
		lended: 0,
	}

	library.books["The Little Go Book"] = BookEntry{
		total:  2,
		lended: 0,
	}

	library.books["Go Bootcamp"] = BookEntry{
		total:  1,
		lended: 0,
	}
	//  - Add at least 3 members...
	library.members["John Doe"] = Member{
		name:  "John Doe",
		books: make(map[Title]LendAudit),
	}

	library.members["Jane Doe"] = Member{
		name:  "Jane Doe",
		books: make(map[Title]LendAudit),
	}

	library.members["Bob House"] = Member{
		name:  "Bob House",
		books: make(map[Title]LendAudit),
	}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMembersAudits(&library)

	//  - Check out a book...

	johnDoe := library.members["John Doe"]
	janeDoe := library.members["Jane Doe"]
	bobHouse := library.members["Bob House"]

	checkoutBook(&library, "Webapps in Go", &johnDoe)
	checkoutBook(&library, "Learning Go", &janeDoe)
	checkoutBook(&library, "The Little Go Book", &bobHouse)

	fmt.Println("\nAfter checkout:")
	printLibraryBooks(&library)
	printMembersAudits(&library)

	//  - Check in a book...
	returnBook(&library, "Webapps in Go", &johnDoe)
	returnBook(&library, "Learning Go", &janeDoe)
	returnBook(&library, "The Little Go Book", &bobHouse)

	fmt.Println("\nAfter return:")
	printLibraryBooks(&library)
	printMembersAudits(&library)

}
