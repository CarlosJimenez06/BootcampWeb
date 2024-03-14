package web

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Custom errors
var (
	ErrInvalidContentType = errors.New("invalid content type")
)

// RequestJSON decodifica el body de la petición en v
func RequestJSON(r *http.Request, v interface{}) (err error) {

	// Verifica que el content type sea application/json
	if r.Header.Get("Content-Type") != "application/json" {
		return ErrInvalidContentType
	}

	// Codifica el body en v, de haber un error lo retorna
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return err
	}

	return
}
