package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// структура ответа
type healthCheckResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

// обрабатывает GET
func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := healthCheckResponse{
		Status:      "available",
		ProjectName: "Library API",
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}