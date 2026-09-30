package models

import "fmt"

type Book struct {
    Title    string
    Author   string
    Year     int
    IsIssued bool
    ReaderID *int
}

// описание книги с учётом её статуса
func (b Book) String() string {
    status := "в библиотеке"
    if b.IsIssued && b.ReaderID != nil {
        status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
    }
    return fmt.Sprintf("\"%s\" (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

//  выдаёт книгу читателю
func (b *Book) IssueBook(reader *Reader) {
    if b.IsIssued {
        fmt.Printf("Книга %s уже кому-то выдана\n", b.Title)
        return
    }
    if !reader.IsActive {
        fmt.Printf("Читатель %s %s не активен и не может получить книгу.\n",
            reader.FirstName, reader.LastName)
        return
    }
    b.IsIssued = true
    b.ReaderID = &reader.ID
    fmt.Printf("Книга %s была выдана читателю %s %s\n",
        b.Title, reader.FirstName, reader.LastName)
}

//  возвращает книгу в библиотеку
func (b *Book) ReturnBook() {
    if !b.IsIssued {
        fmt.Printf("Книга %s и так в библиотеке\n", b.Title)
        return
    }
    b.IsIssued = false
    b.ReaderID = nil
    fmt.Printf("Книга %s возвращена в библиотеку\n", b.Title)
}

type Reader struct {
    ID        int
    FirstName string
    LastName  string
    IsActive  bool
}

//  возвращает описание читателя
func (r Reader) String() string {
    status := "активен"
    if !r.IsActive {
        status = "не активен"
    }
    return fmt.Sprintf("Читатель #%d: %s %s (%s)", r.ID, r.FirstName, r.LastName, status)
}

func (r *Reader) AssignBook(b *Book) {
    fmt.Printf("Читатель %s %s взял книгу %s\n", r.FirstName, r.LastName, b)
}

// делает читателя неактивным
func (r *Reader) Deactivate() {
    r.IsActive = false
}

type Library struct {
    Books   map[int]*Book
    Readers map[int]*Reader
}