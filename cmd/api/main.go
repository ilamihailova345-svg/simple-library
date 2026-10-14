package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type healthCheckResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

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


type infoResponse struct {
	Version    string `json:"version"`
	ServerTime string `json:"server_time"`
}

func infoHandler(w http.ResponseWriter, r *http.Request) {
	response := infoResponse{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
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


func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusNotFound) // 404
	w.Write([]byte("This page does not exist"))
}


func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler)
	http.HandleFunc("/api/v1/info", infoHandler)
	http.HandleFunc("/notfound", notFoundHandler)

	fmt.Println("Сервер запущен на http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}