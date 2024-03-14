package handlers

import (
	"ejemplo1RepoWay/app/internal"
	"ejemplo1RepoWay/app/platform/tools"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/bootcamp-go/web/request"
	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

// Es un constructor que devuelve un nuevo DefaultTask.
func NewDefaultTask(sv internal.TaskService) *DefaultTask {
	return &DefaultTask{
		service: sv,
	}
}

// Inyección de dependencias
type DefaultTask struct {
	service internal.TaskService
}

// Usado para enviar una tarea.
type TaskJSON struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// Usado para guardar una tarea.
type TaskRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

func (d *DefaultTask) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}
		// - parse to map (dynamic)
		bodyMap := make(map[string]any)
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}
		// - validate fields
		if err := tools.CheckField(bodyMap, "title", "description", "Done"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": fmt.Sprintf("%s is required", fieldError.Field)})
				return
			}
			response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			return
		}
		// - parse JSON to struct
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid request body"})
			return
		}
		// - validate the task
		if body.Title == "" || len(body.Title) > 25 {
			response.JSON(w, http.StatusBadRequest, map[string]any{"message": "Title is required and must be less than 25 characters"})
			return
		}

		// PROCESS
		// - serialize the request body into a Task
		task := internal.Task{
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}
		// - save the task, si no se puede guardar, devolver este switch de errores
		if err := d.service.Save(&task); err != nil {
			switch {
			case errors.Is(err, internal.ErrorTaskDuplicated):
				response.JSON(w, http.StatusConflict, map[string]any{"message": "task already exists"})
			case errors.Is(err, internal.ErrorTaskInvalidField):
				response.JSON(w, http.StatusBadRequest, map[string]any{"message": "invalid task"})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// RESPONSE
		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}

		response.JSON(w, http.StatusCreated, map[string]any{
			"message": "task created",
			"data":    data,
		})

	}
}

func (d *DefaultTask) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - parse the id from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
		}

		// PROCESS
		// - delete the task
		if err := d.service.Delete(id); err != nil {
			switch {
			case errors.Is(err, internal.ErrorTaskNotFound):
				response.JSON(w, http.StatusNotFound, map[string]any{"message": "task not found"})
			default:
				response.JSON(w, http.StatusInternalServerError, map[string]any{"message": "internal server error"})
			}
			return
		}

		// RESPONSE
		w.WriteHeader(http.StatusNoContent)
		response.Text(w, http.StatusOK, "task deleted")
	}
}

func (d *DefaultTask) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - parse the id from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}
		// - read into bytes
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - parse to map (dynamic)
		bodyMap := make(map[string]any)
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - validate fields
		if err := tools.CheckField(bodyMap, "title", "description", "Done"); err != nil {
			var fieldError *tools.FieldError
			if errors.As(err, &fieldError) {
				response.Text(w, http.StatusBadRequest, fmt.Sprintf("%s is required", fieldError.Field))
				return
			}
			response.Text(w, http.StatusInternalServerError, "internal server error")
			return
		}
		// - parse JSON to struct
		var body TaskRequestBody
		if err := json.Unmarshal(bytes, &body); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - validate the task
		if body.Title == "" || len(body.Title) > 25 {
			response.Text(w, http.StatusBadRequest, "Title is required and must be less than 25 characters")
			return
		}

		// PROCESS
		// - serialize the request body into a Task
		task := internal.Task{
			ID:          id,
			Title:       body.Title,
			Description: body.Description,
			Done:        body.Done,
		}
		// - update the task
		if err := d.service.Update(task); err != nil {
			switch {
			case errors.Is(err, internal.ErrorTaskNotFound):
				response.Text(w, http.StatusNotFound, "task not found")
			case errors.Is(err, internal.ErrorTaskInvalidField):
				response.Text(w, http.StatusBadRequest, "invalid task")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}

		// RESPONSE
		data := TaskJSON{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Done:        task.Done,
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "task updated",
			"data":    data,
		})
	}
}

func (d *DefaultTask) UpdatePartial() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// REQUEST
		// - parse the id from the URL
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			response.Text(w, http.StatusBadRequest, "invalid id")
			return
		}
		// - parse to map (dynamic)
		bodyMap := make(map[string]any)
		if err := request.JSON(r, &bodyMap); err != nil {
			response.Text(w, http.StatusBadRequest, "invalid request body")
			return
		}
		// - validate fields
		if title, ok := bodyMap["title"]; ok {
			titleString, ok := title.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "invalid title")
				return
			}
			if titleString == "" || len(titleString) > 25 {
				response.Text(w, http.StatusBadRequest, "Title is required and must be less than 25 characters")
				return
			}
		}
		if description, ok := bodyMap["description"]; ok {
			descriptionString, ok := description.(string)
			if !ok {
				response.Text(w, http.StatusBadRequest, "invalid description")
				return
			}
			if descriptionString == "" || len(descriptionString) > 250 {
				response.Text(w, http.StatusBadRequest, "Description is required and must be less than 250 characters")
				return
			}
		}

		// PROCESS
		// - update partially the task
		if err := d.service.UpdatePartial(id, bodyMap); err != nil {
			switch {
			case errors.Is(err, internal.ErrorTaskNotFound):
				response.Text(w, http.StatusNotFound, "task not found")
			case errors.Is(err, internal.ErrorTaskInvalidField):
				response.Text(w, http.StatusBadRequest, "invalid task")
			case errors.Is(err, internal.ErrorTaskDuplicated):
				response.Text(w, http.StatusConflict, "task already exists")
			default:
				response.Text(w, http.StatusInternalServerError, "internal server error")
			}
			return
		}
	}
}
