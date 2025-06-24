package main

import (
	"log"
	"net/http"

	"github.com/Handruka/workmate_task_go.git/internal/api"
	"github.com/Handruka/workmate_task_go.git/internal/service"
	"github.com/Handruka/workmate_task_go.git/internal/storage"
	"github.com/gorilla/mux"
)

func main() {
	// 1. Инициализация хранилища и сервиса
	store := storage.New()
	taskService := service.NewTaskService(store)
	handler := api.NewHandler(taskService)

	// 2. Настройка роутера
	router := mux.NewRouter()
	handler.RegisterRoutes(router)

	// 3. Запуск сервера
	addr := ":8080"
	log.Printf("Server is running on %s...", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
