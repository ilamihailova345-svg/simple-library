package main

import (
	"fmt"

	"library-app/config"
	"library-app/library"
	"library-app/notifications"
)

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	fmt.Println("\n__ Тест config __")

	configWithPort := map[string]string{"PORT": "8080", "HOST": "localhost"}
	configWithoutPort := map[string]string{"HOST": "localhost"}

	port, err := config.GetPortFromConfig(configWithPort)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт найден:", port)
	}

	port, err = config.GetPortFromConfig(configWithoutPort)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт найден:", port)
	}

	fmt.Println("\n__ Библиотека __")

	myLibrary := library.New()

	book1, err := myLibrary.AddBook("Преступление и наказание", "Достоевский", 1866)
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
		return
	}

	book2, err := myLibrary.AddBook("Война и мир", "Толстой", 1869)
	if err != nil {
		fmt.Println("Ошибка добавления книги:", err)
		return
	}

	_, err = myLibrary.AddBook("Преступление и наказание", "Достоевский", 1866)
	if err != nil {
		fmt.Println("Ожидаемая ошибка при добавлении дубликата:", err)
	}

	user1 := myLibrary.AddReader("Шлепок", "Петров")
	user2 := myLibrary.AddReader("Газан", "Иванов")

	fmt.Println("Книг:", len(myLibrary.Books), "Читателей:", len(myLibrary.Readers))


	fmt.Println("\n1. Успешная выдача книги")
	err = myLibrary.IssueBookToReader(book1.ID, user1.ID)
	if err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана.")
	}
	fmt.Println("Статус книги:", book1)

	fmt.Println("\n2. Попытка выдать уже выданную книгу")
	err = myLibrary.IssueBookToReader(book1.ID, user2.ID)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	fmt.Println("\n3. Попытка выдать книгу несуществующему читателю")
	err = myLibrary.IssueBookToReader(book2.ID, 999)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	// === Сценарий 4 ===
	fmt.Println("\n4. Успешный возврат книги")
	err = myLibrary.ReturnBook(book1.ID)
	if err != nil {
		fmt.Println("Ошибка возврата:", err)
	} else {
		fmt.Println("Книга успешно возвращена.")
	}
	fmt.Println("Статус книги:", book1)

	fmt.Println("\n5. Попытка вернуть книгу, которая уже в библиотеке")
	err = myLibrary.ReturnBook(book1.ID)
	if err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	fmt.Println("\n__ Уведомления __")
	emailNotifier := notifications.EmailNotifier{EmailAddress: "student@example.com"}
	smsNotifier := notifications.SMSNotifier{PhoneNumber: "+79991234567"}

	notifiers := []notifications.Notifier{emailNotifier, smsNotifier}
	for _, n := range notifiers {
		n.Notify("Ваша книга просрочена!")
	}

	fmt.Println("\n__ Все книги в библиотеке __")
	for _, b := range myLibrary.GetAllBooks() {
		fmt.Println(b)
	}
}