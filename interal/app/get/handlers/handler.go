package handlers

import (
	"net/http"
)

type Usecase interface {
	Get(uid string) string
}
type Handler struct {
	Usecase Usecase
}

func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	hasUid := r.URL.Query().Has("uid")
	if hasUid {
		uid := r.URL.Query().Get("uid")
		w.Write([]byte(h.Usecase.Get(uid)))
	} else {
		w.Write([]byte("Nothing"))
	}
}
