package users

import (
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", usersHome)
	mux.HandleFunc("/test", test)
}

func usersHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User homepage"))
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing"))
}
