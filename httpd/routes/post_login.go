package routes

import (
	"encoding/json"
	"gradientApi/httpd/models/requests"
	"gradientApi/httpd/models/responses"
	"net/http"
)

func LoginPost(w http.ResponseWriter, r *http.Request) {
	var c requests.Credentials
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		responses.WriteError(w, "invalid request", http.StatusUnprocessableEntity)
		return
	}
	valid, err := c.Validate()
	if err != nil{
		responses.WriteError(w, "server error: " + err.Error(), http.StatusInternalServerError)
	}
	if !valid{
		responses.WriteError(w, "check credentials format", http.StatusUnprocessableEntity)
	}
	auth, err := c.Authenticate()
	if err != nil{
		responses.WriteError(w, "server error: " + err.Error(), http.StatusInternalServerError)
	}
	if !auth {
		responses.WriteError(w, "username and password dont match", http.StatusUnauthorized)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}