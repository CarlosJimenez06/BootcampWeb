package web

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, status int, v interface{}) {
	// Codifica los headers y el status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Verifica si hay un error al codificar el JSON, recodifica los headers y el status, envía el mensaje de error
	if err := json.NewEncoder(w).Encode(v); err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}
}
