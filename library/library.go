package library

import (
	"fmt"

	"library-app/domain"
)

type Library struct {
	Books        []*domain.Book
	Readers      []*domain.Reader
	lastBookID   int
	lastReaderID int
}

//  создаёт новую пустую библиотеку
func New() *Library {
	return &Library{
		Books:   []*domain.Book{},
		Readers: []*domain.Reader{},
	}
}

// добавляет книгу. Ошибка, если дубликат
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

//  добавляет читателя
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

//  книгу по ID
func (lib *Library) FindBookByID(id int) (*domain.Book, error) {
	for _, b := range lib.Books {
		if b.ID == id {
			return b, nil
		}
	}
	return nil, fmt.Errorf("книга с ID %d не найдена", id)
}

//  читателя по ID
func (lib *Library) FindReaderByID(id int) (*domain.Reader, error) {
	for _, r := range lib.Readers {
		if r.ID == id {
			return r, nil
		}
	}
	return nil, fmt.Errorf("читатель с ID %d не найден", id)
}

//   книгу читателю по их ID
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

//  возвращает книгу в библиотеку по её ID
func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}
	return book.ReturnBook()
}

//  всех книг
func (lib *Library) GetAllBooks() []*domain.Book {
	return lib.Books
}





// находит максимальные ID среди книг и читателей
// и устанавливает внутренние счетчики библиотеки на эти значения
// необходимо вызывать после загрузки данных из файла
func (lib *Library) UpdateIDs() {
	maxBookID := 0
	for _, book := range lib.Books {
		if book.ID > maxBookID {
			maxBookID = book.ID
		}
	}
	lib.lastBookID = maxBookID

	maxReaderID := 0
	for _, reader := range lib.Readers {
		if reader.ID > maxReaderID {
			maxReaderID = reader.ID
		}
	}
	lib.lastReaderID = maxReaderID
}