package service

import (
	"time"

	"github.com/Handruka/workmate_task_go.git/internal/model"
	"github.com/Handruka/workmate_task_go.git/internal/storage"
	"github.com/testcontainers/testcontainers-go/log"
)

type TaskService struct {
	store *storage.TaskStorage
}

func NewTaskService(store *storage.TaskStorage) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) StartNewTask() *model.Task {
	task := s.store.Create()

	go func(taskID string) {
		s.store.Update(taskID, "выполняется", "")
		log.Printf("Task %s started", taskID)

		time.Sleep(5 * time.Second)

		s.store.Update(taskID, "готово", "Задача выполнена успешно")
		log.Printf("Task %s выполнена", taskID)

	}(task.ID)

	return task
}
