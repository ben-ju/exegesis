package service

import (
	"net/http"

	"github.com/ben-ju/exegesis/internal/repository"
	"github.com/ben-ju/exegesis/internal/router"
)

type BookRoutable struct {
	repo *repository.BookRepository
}

func NewBookRoutable(repo *repository.BookRepository) *BookRoutable {
	return &BookRoutable{repo: repo}
}

func (br *BookRoutable) RegisterRoutes(r *router.Router) {
	r.HandleFunc("/versions", br.listVersions)
}

func (br *BookRoutable) listVersions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Versions"))
	// versions, err := br.repo.ListVersions()
	// if err != nil {
	// 	http.Error(w, "Error fetching versions", http.StatusInternalServerError)
	// 	return
	// }
	// json.NewEncoder(w).Encode(versions)
}
