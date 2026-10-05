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

// возвращает описание книги с учётом её статуса
func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("\"%s\" (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

// выдаёт книгу читателю
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

// возвращает книгу в библиотеку
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

// возвращает описание читателя
func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "не активен"
	}
	return fmt.Sprintf("Читатель #%d: %s %s (%s)", r.ID, r.FirstName, r.LastName, status)
}

// читатель берёт книгу
func (r *Reader) AssignBook(b *Book) {
	fmt.Printf("Читатель %s %s взял книгу %s\n", r.FirstName, r.LastName, b)
}

// делает читателя неактивным
func (r *Reader) Deactivate() {
	r.IsActive = false
}

// делает читателя активным
func (r *Reader) Activate() {
	r.IsActive = true
}


type Library struct {
	Books   map[int]*Book
	Readers map[int]*Reader
}

// ищет книгу по ID 
func (lib *Library) FindBookByID(id int) (*Book, error) {
	for _, book := range lib.Books {
		if book.ID == id {
			return book, nil
		}
	}
	return nil, fmt.Errorf("Книга с этим id %d не найдена в библиотеке!", id)
}

// ищет читателя по ID
func (lib *Library) FindReaderByID(id int) (*Reader, error) {
	for _, reader := range lib.Readers {
		if reader.ID == id {
			return reader, nil
		}
	}
	return nil, fmt.Errorf("Читатель с этим id %d не найден в библиотеке!", id)
}

// выдаёт книгу читателю по их ID
func (lib *Library) IssueBookToReader(bookID int, readerID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	reader, err := lib.FindReaderByID(readerID)
	if err != nil {
		return err
	}

	if err := book.IssueBook(reader); err != nil {
		return err
	}

	return nil
}

// возвращает книгу в библиотеку по её ID
func (lib *Library) ReturnBook(bookID int) error {
	book, err := lib.FindBookByID(bookID)
	if err != nil {
		return err
	}

	if err := book.ReturnBook(); err != nil {
		return err
	}

	return nil
}

// выводит все книги библиотеки в консоль
func (lib *Library) ListAllBooks() {
	for _, book := range lib.Books {
		fmt.Println(book)
	}
}