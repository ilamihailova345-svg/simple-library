package models

import "fmt"

type Book struct {
	ID       int
	Title    string
	Author   string
	Year     int
	IsIssued bool
	ReaderID *int
}

func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("\"%s\" (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

func (b *Book) IssueBook(reader *Reader) error {
	if b.IsIssued {
		return fmt.Errorf("книга '%s' уже кому-то выдана", b.Title)
	}
	if !reader.IsActive {
		return fmt.Errorf("читатель %s %s не активен и не может получить книгу",
			reader.FirstName, reader.LastName)
	}
	b.IsIssued = true
	b.ReaderID = &reader.ID
	return nil
}

func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}

type Reader struct {
	ID        int
	FirstName string
	LastName  string
	IsActive  bool
}

func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "не активен"
	}
	return fmt.Sprintf("Читатель #%d: %s %s (%s)", r.ID, r.FirstName, r.LastName, status)
}

func (r *Reader) Deactivate() {
	r.IsActive = false
}

func (r *Reader) Activate() {
	r.IsActive = true
}

type Library struct {
	Books        map[int]*Book
	Readers      map[int]*Reader
	lastBookID   int
	lastReaderID int
}

// новая пустая библиотека
func New() *Library {
	return &Library{
		Books:   make(map[int]*Book),
		Readers: make(map[int]*Reader),
	}
}

func (lib *Library) AddBook(title, author string, year int) (*Book, error) {
	for _, b := range lib.Books {
		if b.Title == title && b.Author == author {
			return nil, fmt.Errorf("книга '%s' (%s) уже есть в библиотеке", title, author)
		}
	}
	lib.lastBookID++
	book := &Book{
		ID:     lib.lastBookID,
		Title:  title,
		Author: author,
		Year:   year,
	}
	lib.Books[book.ID] = book
	return book, nil
}

// добавляет читателя
func (lib *Library) AddReader(firstName, lastName string) *Reader {
	lib.lastReaderID++
	reader := &Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	lib.Readers[reader.ID] = reader
	return reader
}

func (lib *Library) FindBookByID(id int) (*Book, error) {
	book, ok := lib.Books[id]
	if !ok {
		return nil, fmt.Errorf("книга с ID %d не найдена", id)
	}
	return book, nil
}

func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	reader, ok := lib.Readers[id]
	if !ok {
		return nil, fmt.Errorf("читатель с ID %d не найден", id)
	}
	return reader, nil
}

func (lib *Library) IssueBookToReader(bookID, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}
	return book.IssueBook(reader)
}

func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

func (lib *Library) ListAllBooks() []*Book {
	result := make([]*Book, 0, len(lib.Books))
	for _, b := range lib.Books {
		result = append(result, b)
	}
	return result
}