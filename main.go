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

    lib := models.Library{
        Books:   map[int]*models.Book{1: &book1},
        Readers: map[int]*models.Reader{1: &user1},
    }
    fmt.Println("Книг в библиотеке:", len(lib.Books))
    fmt.Println("Читателей:", len(lib.Readers))

    // --- Система уведомлений ---
    emailNotifier := models.EmailNotifier{EmailAddress: "student@example.com"}
    smsNotifier := models.SMSNotifier{PhoneNumber: "+79991234567"}

    notifiers := []models.Notifier{emailNotifier, smsNotifier}

    for _, n := range notifiers {
        n.Notify("Ваша книга просрочена!")
    }
}