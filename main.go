package main

import (
	"fmt"

	"simple-library/models"
)

func main() {
	book := models.Book{
		Title:  "Мастер и Маргарита",
		Author: "Булгаков",
		Year:   1967,
	}

	fmt.Println(book)

	book.IssueBook()
	fmt.Println(book)

	book.IssueBook()

	book.ReturnBook()
	fmt.Println(book)

	book.ReturnBook()
}