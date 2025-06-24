package storage_test

import (
	"testing"
	"time"

	"github.com/Handruka/workmate_task_go.git/internal/storage"
	"github.com/stretchr/testify/assert"
)

func TestTaskStorage_CreateAndGet(t *testing.T) {
	store := storage.New()

	task := store.Create()

	assert.NotEmpty(t, task.ID, "Ожидался непустой ID задачи")

	got, ok := store.Get(task.ID)
	assert.True(t, ok, "Задача должна быть найдена по ID")
	assert.Equal(t, task.ID, got.ID, "ID полученной задачи не совпадает с ожидаемым")
}

func TestTaskStorage_Update(t *testing.T) {
	store := storage.New()
	task := store.Create()

	ok := store.Update(task.ID, "выполняется", "")
	assert.True(t, ok, "Ожидалось успешное обновление задачи")

	updated, _ := store.Get(task.ID)
	assert.Equal(t, "выполняется", updated.Status, "Статус задачи не обновился")
	assert.WithinDuration(t, time.Now(), updated.UpdatedAT, time.Second, "Время обновления задачи некорректное")
}

func TestTaskStorage_Delete(t *testing.T) {
	store := storage.New()
	task := store.Create()

	ok := store.Delete(task.ID)
	assert.True(t, ok, "Ожидалось успешное удаление задачи")

	_, exists := store.Get(task.ID)
	assert.False(t, exists, "Задача не должна существовать после удаления")
}

func TestTaskStorage_GetAll(t *testing.T) {
	store := storage.New()

	store.Create()
	store.Create()
	store.Create()

	tasks := store.GetAll()
	assert.Len(t, tasks, 3, "Ожидалось получить 3 задачи")
}
