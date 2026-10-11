package main

import "fmt"

type Book struct {
    Title  string
    Author string
    Year   int
}

// возвращает строку с описанием книги
func (b Book) String() string {
    return fmt.Sprintf(`"%s" (%s, %d)`, b.Title, b.Author, b.Year)
}