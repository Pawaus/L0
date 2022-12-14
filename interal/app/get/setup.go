package get

import (
	"api/interal/app/get/handlers"
	"api/interal/app/get/repo"
	"api/interal/app/get/usecase"
	"api/interal/pkg/cache"
	"database/sql"
)

func Setup(db *sql.DB, ch *cache.Cache) handlers.Handler {
	r := repo.NewRepo(db, ch)
	uc := usecase.Usecase{Repo: r}
	handler := handlers.Handler{Usecase: uc}
	return handler
}
