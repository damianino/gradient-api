package responses

import (
	"encoding/json"
	"net/http"
)

type InternalServerError struct {
}

type InvalidRequest struct {
}

type Unauthorized struct {
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Error(s string) *ErrorResponse{
	r := ErrorResponse{
		Error: s,
	}
	return &r
}

func (r *ErrorResponse)Bytes() []byte {
	b, _ := json.Marshal(r)
	return b
}

func WriteError(w http.ResponseWriter, s string, status int){
	b := Error(s).Bytes()
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write(b)
}