package main

import (
	"fmt"

	"simple-library/models"
)

func main() {


	fmt.Println("\n Test1 - config")
	configWithPort := map[string]string{
		"PORT": "8080",
		"HOST": "localhost",
	}
	configWithoutPort := map[string]string{
		"HOST": "localhost",
	}

	port, err := GetPortFromConfig(configWithPort)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт найден:", port)
	}

	port, err = GetPortFromConfig(configWithoutPort)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Порт найден:", port)
	}



	fmt.Println("\n__ Библиотека __")
	fmt.Println("Проект запущен.")

	book1 := models.Book{
		ID:     1,
		Title:  "Преступление и наказание",
		Author: "Достоевский",
		Year:   1866,
	}

	user1 := models.Reader{
		ID:        1,
		FirstName: "Шлепок",
		LastName:  "Обувние",
		IsActive:  true,
	}
	reader2 := models.Reader{
		ID:        2,
		FirstName: "Akaki",
		LastName:  "Kakievich",
		IsActive:  true,
	}


	// выдаём книгу user1
	if err := book1.IssueBook(&user1); err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга выдана:", book1)
	}

	// ОШИБКА: выдать ту же книгу reader2
	if err := book1.IssueBook(&reader2); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}



	user1.Deactivate()
	fmt.Println(" Test2 - user1")

	// возврат книги
	if err := book1.ReturnBook(); err != nil {
		fmt.Println("Ошибка возврата:", err)
	} else {
		fmt.Println("Книга возвращена:", book1)
	}

	// ОШИБКА: повторный возврат
	if err := book1.ReturnBook(); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	myLibrary := models.Library{
		Books:   map[int]*models.Book{1: &book1},
		Readers: map[int]*models.Reader{1: &user1},
	}
	fmt.Println("Книг в библиотеке:", len(myLibrary.Books))
	fmt.Println("Читателей:", len(myLibrary.Readers))



	fmt.Println(" __ Система уведомлений __")
	emailNotifier := models.EmailNotifier{EmailAddress: "student@example.com"}
	smsNotifier := models.SMSNotifier{PhoneNumber: "+79991234567"}

	notifiers := []models.Notifier{emailNotifier, smsNotifier}
	for _, n := range notifiers {
		n.Notify("Ваша книга просрочена!")
	}



	fmt.Println("\n__ Test3 - выдача книг __")

	user1.Activate()

	if err := myLibrary.IssueBookToReader(1, 1); err != nil {
		fmt.Println("Ошибка выдачи:", err)
	} else {
		fmt.Println("Книга успешно выдана читателю.")
	}

	book, _ := myLibrary.FindBookByID(1)
	if book != nil {
		fmt.Println("Статус книги после выдачи:", book)
	}

	if err := myLibrary.IssueBookToReader(99, 1); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	
	
	fmt.Println("\n__ Тестируем возврат книги через Library.ReturnBook __")

	if err := myLibrary.ReturnBook(1); err != nil {
		fmt.Println("Ошибка возврата:", err)
	} else {
		fmt.Println("Книга успешно возвращена через библиотеку.")
	}

	book, _ = myLibrary.FindBookByID(1)
	if book != nil {
		fmt.Println("Статус книги после возврата:", book)
	}

	// ОШИБКА: повторный возврат - "книга и так в библиотеке"
	if err := myLibrary.ReturnBook(1); err != nil {
		fmt.Println("Ожидаемая ошибка:", err)
	}

	fmt.Println("\n__ Все книги в библиотеке __")
	myLibrary.ListAllBooks()
}