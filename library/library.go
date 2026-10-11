package library

import (
	"fmt"

	"library-app/domain"
	"library-app/storage"
)

type Library struct {
	Books        []*domain.Book
	Readers      []*domain.Reader
	lastBookID   int
	lastReaderID int
}

// New создаёт новую пустую библиотеку.
func New() *Library {
	return &Library{
		Books:   []*domain.Book{},
		Readers: []*domain.Reader{},
	}
}

// добавляет книгу
func (lib *Library) AddBook(title, author string, year int) (*domain.Book, error) {
	for _, b := range lib.Books {
		if b.Title == title && b.Author == author {
			return nil, fmt.Errorf("книга '%s' (%s) уже есть в библиотеке", title, author)
		}
	}
	lib.lastBookID++
	book := &domain.Book{
		ID:     lib.lastBookID,
		Title:  title,
		Author: author,
		Year:   year,
	}
	lib.Books = append(lib.Books, book)
	return book, nil
}

func (lib *Library) AddReader(firstName, lastName string) *domain.Reader {
	lib.lastReaderID++
	reader := &domain.Reader{
		ID:        lib.lastReaderID,
		FirstName: firstName,
		LastName:  lastName,
		IsActive:  true,
	}
	lib.Readers = append(lib.Readers, reader)
	return reader
}

func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, b := range lib.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
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

func (lib *Library) GetAllBooks() []*domain.Book {
	return lib.Books
}

// сохраняет книги библиотеки в CSV
func (lib *Library) SaveToCSV(filename string) error {
	return storage.SaveBooksToCSV(lib.Books, filename)
}

// загружает книги из CSV, полностью заменяя текущий список
func (lib *Library) LoadFromCSV(filename string) error {
	books, err := storage.LoadBooksFromCSV(filename)
	if err != nil {
		return err
	}
	lib.Books = books

	maxID := 0
	for _, b := range books {
		if b.ID > maxID {
			maxID = b.ID
		}
	}
	lib.lastBookID = maxID
	return nil
}