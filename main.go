package main

import (
	"fmt"
	"os"

	"library-app/library"
	"library-app/notifications"
)

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	myLibrary := library.New()

	_, err := myLibrary.AddBook("Преступление и наказание", "Достоевский", 1866)
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
		return
	}

	_, err = myLibrary.AddBook("Война и мир", "Толстой", 1869)
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
		return
	}

	user1 := myLibrary.AddReader("Колян", "Вованыч")
	_ = user1

	if len(os.Args) >= 3 {
		command := os.Args[1]
		filename := os.Args[2]

		switch command {
		case "save":
			err := myLibrary.SaveToCSV(filename)
			if err != nil {
				fmt.Println("Ошибка сохранения:", err)
				os.Exit(1)
			}
			fmt.Printf("Книги сохранены в %s\n", filename)
			return

		case "load":
			err := myLibrary.LoadFromCSV(filename)
			if err != nil {
				fmt.Println("Ошибка загрузки:", err)
				os.Exit(1)
			}
			fmt.Printf("Книги загружены из %s:\n", filename)
			for _, b := range myLibrary.GetAllBooks() {
				fmt.Println(b)
			}
			return

		default:
			fmt.Printf("Неизвестная команда: %s\n", command)
			fmt.Println("Доступные команды:")
			fmt.Println("  save <имя_файла.csv>")
			fmt.Println("  load <имя_файла.csv>")
			os.Exit(1)
		}
	}

	fmt.Println("\nДемо-режим. Доступные команды:")
	fmt.Println("  go run . save books.csv")
	fmt.Println("  go run . load books.csv")
	fmt.Println()

	fmt.Println("Книги в библиотеке:")
	for _, b := range myLibrary.GetAllBooks() {
		fmt.Println(b)
	}
	fmt.Println()

	fmt.Println("Уведомления:")
	emailNotifier := notifications.EmailNotifier{EmailAddress: "student@example.com"}
	smsNotifier := notifications.SMSNotifier{PhoneNumber: "+79991234567"}

	notifiers := []notifications.Notifier{emailNotifier, smsNotifier}
	for _, n := range notifiers {
		n.Notify("Ваша книга просрочена!")
	}
}