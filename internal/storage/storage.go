package storage

import (
	"sync"
	"time"

	"github.com/Handruka/workmate_task_go.git/internal/model"
	"github.com/google/uuid"
)

type TaskStorage struct {
	mu    sync.RWMutex
	Tasks map[string]*model.Task
}

func New() *TaskStorage {
	res := TaskStorage{
		Tasks: make(map[string]*model.Task),
	}
	return &res
}

func (s *TaskStorage) Create() *model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.NewString()
	now := time.Now()

	task := &model.Task{
		ID:        id,
		Status:    "Задача создана",
		CreatedAT: now,
		UpdatedAT: now,
	}
	s.Tasks[id] = task
	return task
}

func (s *TaskStorage) Get(id string) (*model.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	task, exists := s.Tasks[id]
	return task, exists
}

func (s *TaskStorage) Update(id, status, result string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exist := s.Tasks[id]
	if !exist {
		return false
	}

	task.Status = status
	task.UpdatedAT = time.Now()
	if result != "" {
		task.Result = result
	}

	return true
}

func (s *TaskStorage) Delete(id string) bool {
	if _, exist := s.Tasks[id]; exist {
		delete(s.Tasks, id)
		return true
	}
	return false
}

func (s *TaskStorage) GetAll() []*model.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*model.Task, 0, len(s.Tasks))
	for _, task := range s.Tasks {
		tasks = append(tasks, task)
	}
	return tasks
}
