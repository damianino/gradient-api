package routes

import (
	"encoding/json"
	"gradientApi/httpd/models/responses"
	"gradientApi/httpd/models/requests"
	"gradientApi/httpd/models/db"
	"net/http"
)

func RegisterPost(w http.ResponseWriter, r *http.Request) {
	var c requests.Credentials
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		responses.WriteError(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}
	var newUser dbm.User
	newUser.Username = c.Username
	newUser.Password = c.Password

	err = newUser.Insert()
	if err != nil {
		responses.WriteError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}