package update

import (
	"api/interal/app/update/handlers"
	"api/interal/app/update/repo"
	"api/interal/app/update/usecase"
	"api/interal/pkg/cache"
	"database/sql"
)

func Setup(db *sql.DB, ch *cache.Cache) *handlers.Handler {
	r := repo.NewRepo(db, ch)
	uc := usecase.NewUseCase(r)
	handler := handlers.NewHandler(uc)
	return handler
}
