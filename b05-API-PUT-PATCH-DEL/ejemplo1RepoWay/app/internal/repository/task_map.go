package repository

import (
	"ejemplo1RepoWay/app/internal"
)

// TaskMap es un repositorio de tareas en memoria que usa un map como base de datos.
type TaskMap struct {
	// DB es un map que contiene las tareas.
	// - key: ID de la tarea.
	// - value: Tarea. (structured data)
	db     map[int]internal.Task
	lastID int
}

// Crea una nueva instancia de TaskMap. Actua como un constructor.
func NewTaskMap(db map[int]internal.Task, lastID int) *TaskMap {
	// Default values
	// - De no haber tareas, se crea un map vacío. El id de la última tarea es 0.
	if db == nil {
		db = make(map[int]internal.Task)
	}

	// Retorna una nueva instancia de TaskMap con los valores por proporcionados
	return &TaskMap{
		db:     db,
		lastID: lastID,
	}
}

// Save guarda una tarea en la base de datos.
func (t *TaskMap) Save(task *internal.Task) (err error) {
	for _, t := range (*t).db {
		if t.Title == task.Title {
			return internal.ErrorTaskDuplicated
		}
	}

	// Increment the lastID and assign it to the task ID.
	(*t).lastID++

	// Set the task ID.
	(*task).ID = (*t).lastID

	// Save the task in the database.
	t.db[(*task).ID] = *task

	return
}

func (t *TaskMap) Update(task internal.Task) (err error) {
	// Check if the task exists.
	if _, ok := (*t).db[task.ID]; !ok {
		return internal.ErrorTaskNotFound
	}

	// Check if the task already exists.
	for _, t := range (*t).db {
		if t.Title == task.Title && t.ID != task.ID {
			return internal.ErrorTaskDuplicated
		}
	}

	(*t).db[task.ID] = task

	return
}

func (t *TaskMap) Delete(id int) (err error) {
	// Check if the task exists.
	if _, ok := (*t).db[id]; !ok {
		return internal.ErrorTaskNotFound
	}

	// Delete the task from the database.
	delete((*t).db, id)
	return
}

func (t *TaskMap) UpdatePartial(id int, fields map[string]interface{}) (err error) {
	// Check if the task exists.
	task, ok := (*t).db[id]
	if !ok {
		return internal.ErrorTaskNotFound
	}

	// Update the task fields.
	for field, value := range fields {
		switch field {
		case "Title", "title":
			title, ok := value.(string)
			if !ok {
				return internal.ErrorTaskInvalidField
			}
			// Check if the task already exists.
			for _, t := range (*t).db {
				if t.Title == title && t.ID != id {
					return internal.ErrorTaskDuplicated
				}
			}
			// Update the task title.
			task.Title = title
		case "Description", "description":
			description, ok := value.(string)
			if !ok {
				return internal.ErrorTaskInvalidField
			}
			// Check if the task already exists.
			for _, t := range (*t).db {
				if t.Description == description && t.ID != id {
					return internal.ErrorTaskDuplicated
				}
			}
			// Update the task description.
			task.Description = description
		case "Done", "done":
			task.Done, ok = value.(bool)
			if !ok {
				return internal.ErrorTaskInvalidField
			}
		default:
		}
	}

	// Update the task in the database. (dado el map, se actualiza el valor dependiente de la key)
	(*t).db[id] = task

	return
}
