package api

import (
	"encoding/json"
	"net/http"

	"workmate/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	taskService *service.TaskService
}

func NewHandler(taskService *service.TaskService) *Handler {
	return &Handler{taskService: taskService}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/task", h.createTask).Methods("POST")
	r.HandleFunc("/task/{id}", h.getTask).Methods("GET")
	r.HandleFunc("/task/{id}", h.deleteTask).Methods("DELETE")
}

func (h *Handler) createTask(w http.ResponseWriter, r *http.Request) {
	task := h.taskService.StartNewTask()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) getTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	task, ok := h.taskService.GetByID(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (h *Handler) deleteTask(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	ok := h.taskService.DeleteByID(id)
	if !ok {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
