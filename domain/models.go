package domain

import "fmt"

type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	IsIssued bool   `json:"is_issued"`
	ReaderID *int   `json:"reader_id"`
}

// возвращает описание книги с учётом её статуса.
func (b Book) String() string {
	status := "в библиотеке"
	if b.IsIssued && b.ReaderID != nil {
		status = fmt.Sprintf("на руках у читателя с ID %d", *b.ReaderID)
	}
	return fmt.Sprintf("\"%s\" (%s, %d), статус: %s", b.Title, b.Author, b.Year, status)
}

// выдаёт книгу читателю.
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

// возвращает книгу в библиотеку.
func (b *Book) ReturnBook() error {
	if !b.IsIssued {
		return fmt.Errorf("книга '%s' и так в библиотеке", b.Title)
	}
	b.IsIssued = false
	b.ReaderID = nil
	return nil
}

type Reader struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active"`
}

// возвращает описание читателя
func (r Reader) String() string {
	status := "активен"
	if !r.IsActive {
		status = "не активен"
	}
	return fmt.Sprintf("Читатель #%d: %s %s (%s)", r.ID, r.FirstName, r.LastName, status)
}

// делает читателя неактивным
func (r *Reader) Deactivate() {
	r.IsActive = false
}

// делает читателя активным
func (r *Reader) Activate() {
	r.IsActive = true
}