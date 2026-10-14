package main

import (
	"fmt"
	"log"
	"os"

	"library-app/cli"
	"library-app/library"
	"library-app/storage"
)

func main() {
	fmt.Println("Проект 'Простая библиотека' запущен.")

	const dbFile = "books.json"

	// загружаем библиотеку из файла
	myLibrary, err := storage.LoadLibraryFromJSON(dbFile)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Файл данных не найден, создана новая библиотека.")
			myLibrary = library.New()
		} else {
			log.Fatalf("Ошибка при загрузке библиотеки: %v", err)
		}
	} else {
		fmt.Println("Библиотека успешно загружена из файла.")
	}

	cli.Run(myLibrary, dbFile)
}