package models

import "fmt"

// Notifier — интерфейс для всех типов уведомителей.
type Notifier interface {
    Notify(message string)
}

// EmailNotifier отправляет уведомления по email.
type EmailNotifier struct {
    EmailAddress string
}

func (e EmailNotifier) Notify(message string) {
    fmt.Printf("Отправляю email на %s: \"%s\"\n", e.EmailAddress, message)
}

// SMSNotifier отправляет уведомления через SMS.
type SMSNotifier struct {
    PhoneNumber string
}

func (s SMSNotifier) Notify(message string) {
    fmt.Printf("Отправляю SMS на номер %s: \"%s\"\n", s.PhoneNumber, message)
}
