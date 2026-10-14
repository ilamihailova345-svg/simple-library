package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"library-app/library"
	"library-app/storage"
)

// запускает консольный интерфейс.
func Run(lib *library.Library, dbPath string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n-- Меню --")
		fmt.Println("1 - Показать все книги")
		fmt.Println("2 - Добавить книгу")
		fmt.Println("3 - Добавить читателя")
		fmt.Println("4 - Выдать книгу")
		fmt.Println("5 - Вернуть книгу")
		fmt.Println("0 - Сохранить и выйти")
		fmt.Print("Выбор: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Неверный ввод.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Книги:")
			for _, b := range lib.GetAllBooks() {
				fmt.Println(b)
			}

		case 2:
			fmt.Print("Название: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Автор: ")
			author, _ := reader.ReadString('\n')
			author = strings.TrimSpace(author)

			fmt.Print("Год: ")
			yearStr, _ := reader.ReadString('\n')
			yearStr = strings.TrimSpace(yearStr)
			year, err := strconv.Atoi(yearStr)
			if err != nil {
				fmt.Println("Неверный год.")
				continue
			}

			if _, err := lib.AddBook(title, author, year); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга добавлена.")
			}

		case 3:
			fmt.Print("Имя: ")
			firstName, _ := reader.ReadString('\n')
			firstName = strings.TrimSpace(firstName)

			fmt.Print("Фамилия: ")
			lastName, _ := reader.ReadString('\n')
			lastName = strings.TrimSpace(lastName)

			lib.AddReader(firstName, lastName)
			fmt.Println("Читатель добавлен.")

		case 4:
			fmt.Print("ID книги: ")
			bookIDStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

			fmt.Print("ID читателя: ")
			readerIDStr, _ := reader.ReadString('\n')
			readerID, _ := strconv.Atoi(strings.TrimSpace(readerIDStr))

			if err := lib.IssueBookToReader(bookID, readerID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга выдана.")
			}

		case 5:
			fmt.Print("ID книги: ")
			bookIDStr, _ := reader.ReadString('\n')
			bookID, _ := strconv.Atoi(strings.TrimSpace(bookIDStr))

			if err := lib.ReturnBook(bookID); err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Книга возвращена.")
			}

		case 0:
			fmt.Println("Сохранение данных и выход.")
			if err := storage.SaveLibraryToJSON(dbPath, lib); err != nil {
				fmt.Println("Ошибка сохранения:", err)
			}
			return

		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}