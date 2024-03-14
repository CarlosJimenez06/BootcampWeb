package service

import (
	"ejemplo1RepoWay/app/internal"
)

// Constructor
func NewDefaultTask(rp internal.TaskRepository) *TaskDefault {
	// Crear una nueva instancia de TaskDefault
	return &TaskDefault{
		repository: rp,
	}
}

// Injección de dependencias
type TaskDefault struct {
	repository internal.TaskRepository
}

func (t *TaskDefault) Save(task *internal.Task) (err error) {
	// Guardar la tarea en la base de datos.
	err = (*t).repository.Save(task)
	return
}

func (t *TaskDefault) Update(task internal.Task) (err error) {
	// Actualizar la tarea en la base de datos.
	err = (*t).repository.Update(task)
	return
}

func (t *TaskDefault) Delete(id int) (err error) {
	// Eliminar la tarea de la base de datos.
	err = (*t).repository.Delete(id)
	return
}

func (t *TaskDefault) UpdatePartial(id int, fields map[string]interface{}) (err error) {
	// Actualizar parcialmente la tarea en la base de datos.
	err = (*t).repository.UpdatePartial(id, fields)
	return
}
