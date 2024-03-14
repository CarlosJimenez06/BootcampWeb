package internal

import (
	"errors"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
}

// ERRORS
// - Tarea no encontrada
// - Tarea duplicada
// - Tarea con campos inválidos
var (
	ErrorTaskNotFound     = errors.New("task not found")
	ErrorTaskDuplicated   = errors.New("task already exists")
	ErrorTaskInvalidField = errors.New("task is invalid")
)

// INTERFACES
// - TaskRepository
// - TaskService
type TaskRepository interface {
	Save(task *Task) (err error)
	Update(task Task) (err error)
	Delete(id int) (err error)
	UpdatePartial(id int, fields map[string]interface{}) (err error)
}

type TaskService interface {
	Save(task *Task) (err error)
	Update(task Task) (err error)
	Delete(id int) (err error)
	UpdatePartial(id int, fields map[string]interface{}) (err error)
}
