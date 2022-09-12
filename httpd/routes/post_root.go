package routes

import (
	"encoding/json"
	"gradientApi/httpd/models"
	"gradientApi/httpd/models/responses"
	"net/http"
)

func RootPost(w http.ResponseWriter, r *http.Request) {
	var g models.Gradient
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		responses.WriteError(w, "invalid parameters", http.StatusUnprocessableEntity)
		return
	}

	if !g.Validate() {
		responses.WriteError(w, "invalid parameters", http.StatusUnprocessableEntity)
		return
	}
	b, err := g.ToByteArray()
	if err != nil {
		responses.WriteError(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	w.Header().Set("Content-Type", "image/png; charset=UTF-8")
	w.Write(b)
}
