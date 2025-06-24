package service_test

import (
	"testing"
	"time"

	"github.com/Handruka/workmate_task_go.git/internal/service"
	"github.com/Handruka/workmate_task_go.git/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestStartNewTask_Flow(t *testing.T) {
	store := storage.New()
	svc := service.NewTaskService(store)

	task := svc.StartNewTask()

	assert.Equal(t, "Задача создана", task.Status, "Статус сразу после создания должен быть 'Задача создана'")

	time.Sleep(6 * time.Second)

	updated, ok := store.Get(task.ID)
	assert.True(t, ok, "Задача должна существовать")
	assert.Equal(t, "готово", updated.Status, "Статус должен быть 'готово' после выполнения")
	assert.NotEmpty(t, updated.Result, "Результат задачи должен быть заполнен")
}
