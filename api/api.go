package api

import (
	"encoding/json"
	"net/http"
)

type GetGenericResponse struct {
	Code int
	Msg  string
}

type GetProductByIdResponse struct {
	Code  string
	Id    string
	Desc  string
	Price string
	Date  string
}

type Error struct {
	Code int
	Msg  string
}

func writeError(w http.ResponseWriter, msg string, code int) {
	resp := Error{
		Code: code,
		Msg:  msg,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Ocurred", http.StatusInternalServerError)
	}
	UnauthorizedErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Invalid token", http.StatusUnauthorized)
	}
	PaymentError = func(w http.ResponseWriter, backURL string) {
		writeError(w, backURL, http.StatusBadRequest)
	}
)
