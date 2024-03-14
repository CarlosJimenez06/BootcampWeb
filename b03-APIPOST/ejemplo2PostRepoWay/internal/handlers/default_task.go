package handlers

import (
	"ejemplo2PostRepoWay/internal"
	"ejemplo2PostRepoWay/platform/web"
	"encoding/json"
	"io"
	"net/http"
)

// DefaultTask is a struct that contains the handlers for the tasks
type DefaultTask struct {
	tasks  map[int]internal.Task
	lastID int
}

// Create es un method que crea una nueva task
func (d *DefaultTask) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Validar el token
		token := r.Header.Get("Authorization")
		if token != "12345" {
			web.ResponseJSON(w, http.StatusUnauthorized, map[string]any{"error": "unauthorized"})
			return
		}

		// REQUEST 			(Configura todos lo relacionado con el request)
		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body"})
			return
		}

		// - parse to map (dinamic)
		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body"})
			return
		}

		// - validate fields
		if _, ok := bodyMap["title"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid title"})
			return
		}
		if _, ok := bodyMap["description"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid description"})
			return
		}
		if _, ok := bodyMap["done"]; !ok {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid done"})
			return
		}

		// - parse to struct (static)
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid request body"})
			return
		}

		// PROCESS
		// - Incremento automático del ID
		d.lastID++
		// - Serializar el request body en una task
		task := internal.Task{
			ID:          d.lastID,
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}
		// - Validaciones especificas del task
		if task.Title == "" || len(task.Title) > 100 {
			web.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "invalid title: must be between 1 and 100 characters"})
			return
		}
		// - Agregar el task al map de tasks
		d.tasks[task.ID] = task

		// RESPONSE
		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}
		// - Enviar la task como respuesta
		web.ResponseJSON(w, http.StatusCreated, map[string]any{
			"message": "task created successfully",
			"data":    data,
		})
	}
}

type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type TaskJSON struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Le entra un map de tasks [struct Completo] y un int que es el lastID, retorna un puntero a DefaultTask
func NewDefaultTask(tasks map[int]internal.Task, lastID int) *DefaultTask {
	// Default values
	// Crea un map de tasks [struct Completo]
	defaultTasks := make(map[int]internal.Task)
	defaultLastID := 0

	// If tasks is not nil, assign it to defaultTasks
	if tasks != nil {
		defaultTasks = tasks
	}

	// If lastID is not 0, assign it to defaultLastID
	if lastID != 0 {
		defaultLastID = lastID
	}

	return &DefaultTask{
		tasks:  defaultTasks,
		lastID: defaultLastID,
	}
}
