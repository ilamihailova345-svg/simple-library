package models

import "fmt"

type Book struct {
	Title    string
	Author   string
	Year     int
	IsIssued bool
}

// String возвращает строку с описанием книги в формате
// "Название" (Автор, Год).
func (b Book) String() string {
	return fmt.Sprintf("\"%s\" (%s, %d)", b.Title, b.Author, b.Year)
}

// IssueBook выдаёт книгу читателю.
func (b *Book) IssueBook() {
	if b.IsIssued {
		fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
		return
	}
	b.IsIssued = true
	fmt.Printf("Книга %s была выдана\n", b.Title)
}

// ReturnBook возвращает книгу в библиотеку.
func (b *Book) ReturnBook() {
	if !b.IsIssued {
		fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
		return
	}
	b.IsIssued = false
	fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}