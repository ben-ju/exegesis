package users

import (
	"net/http"

	"github.com/ben-ju/exegesis/internal/router"
)

type UserRoutable struct{}

func (ur *UserRoutable) RegisterRoutes(r *router.Router) {
	r.HandleFunc("/", ur.usersHome)
	r.HandleFunc("/test", ur.test)
}

func (ur *UserRoutable) usersHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User homepage"))
}

func (ur *UserRoutable) test(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Testing"))
}
