package main

import (
    "fmt"

    "simple-library/models"
)

func main() {
    fmt.Println("Проект запущен.")

    book1 := models.Book{
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

    book1.IssueBook(&user1)
    fmt.Println(book1)

    book1.IssueBook(&reader2)

    user1.Deactivate()
    fmt.Println(user1)
    fmt.Println("___")

    book1.ReturnBook()
    book1.IssueBook(&user1)

    user1.AssignBook(&book1)

    myLibrary := models.Library{
        Books:   map[int]*models.Book{1: &book1},
        Readers: map[int]*models.Reader{1: &user1},
    }
    fmt.Println("Книг в библиотеке:", len(myLibrary.Books))
    fmt.Println("Читателей:", len(myLibrary.Readers))

    // Система уведомлений 
    emailNotifier := models.EmailNotifier{EmailAddress: "student@example.com"}
    smsNotifier := models.SMSNotifier{PhoneNumber: "+79991234567"}

    notifiers := []models.Notifier{emailNotifier, smsNotifier}

    for _, n := range notifiers {
        n.Notify("Ваша книга просрочена!")
    }

    fmt.Println("__ Тестируем выдачу книг __")

    err := myLibrary.IssueBookToReader(1, 1)
    if err != nil {
        fmt.Println("Ошибка выдачи:", err)
    }

    book, _ := myLibrary.FindBookByID(1)
    if book != nil {
        fmt.Println("Статус книги после выдачи:", book)
    }

    err = myLibrary.IssueBookToReader(99, 1)
    if err != nil {
        fmt.Println("Ожидаемая ошибка:", err)
    }
}