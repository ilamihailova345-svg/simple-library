package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"library-app/domain"
)

// Storable — интерфейс для типов, которые можно сохранять и загружать.
type Storable interface {
	Save() error
	Load() error
}

// SaveBooksToCSV сохраняет книги в CSV-файл.
func SaveBooksToCSV(books []*domain.Book, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("не удалось создать файл %s: %w", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовок
	header := []string{"ID", "Title", "Author", "Year", "IsIssued", "ReaderID"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("ошибка записи заголовка: %w", err)
	}

	// Данные
	for _, book := range books {
		readerID := ""
		if book.ReaderID != nil {
			readerID = strconv.Itoa(*book.ReaderID)
		}

		record := []string{
			strconv.Itoa(book.ID),
			book.Title,
			book.Author,
			strconv.Itoa(book.Year),
			strconv.FormatBool(book.IsIssued),
			readerID,
		}
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("ошибка записи книги %q: %w", book.Title, err)
		}
	}

	return nil
}

// LoadBooksFromCSV читает книги из CSV-файла.
func LoadBooksFromCSV(filename string) ([]*domain.Book, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("не удалось открыть файл %s: %w", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	if len(records) < 2 {
		return []*domain.Book{}, nil // только заголовок или пусто
	}

	books := make([]*domain.Book, 0, len(records)-1)

	// Пропускаем заголовок
	for _, record := range records[1:] {
		if len(record) < 6 {
			return nil, fmt.Errorf("неверный формат строки: %v", record)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения ID %q: %w", record[0], err)
		}

		year, err := strconv.Atoi(record[3])
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения года %q: %w", record[3], err)
		}

		isIssued, err := strconv.ParseBool(record[4])
		if err != nil {
			return nil, fmt.Errorf("ошибка чтения IsIssued %q: %w", record[4], err)
		}

		var readerID *int
		if record[5] != "" {
			rid, err := strconv.Atoi(record[5])
			if err != nil {
				return nil, fmt.Errorf("ошибка чтения ReaderID %q: %w", record[5], err)
			}
			readerID = &rid
		}

		book := &domain.Book{
			ID:       id,
			Title:    record[1],
			Author:   record[2],
			Year:     year,
			IsIssued: isIssued,
			ReaderID: readerID,
		}
		books = append(books, book)
	}

	return books, nil
}