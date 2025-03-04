package service

import (
	"net/http"
)

func register(r *http.Request) (int, string) {
	if r.Method != "POST" {
		return http.StatusMethodNotAllowed, "invalid method"
	}
	email := r.FormValue("email")
	password := r.FormValue("password")
	if len(email) < 8 || len(password) < 8 {
		return http.StatusBadRequest, "fields are too short"
	}
	return http.StatusOK, "ok"
}
