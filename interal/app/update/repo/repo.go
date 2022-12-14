package repo

import (
	"api/interal/pkg/cache"
	"api/interal/pkg/domain"
	"database/sql"
	"encoding/json"
	"log"
)

type Repo struct {
	db    *sql.DB
	cache *cache.Cache
}

func NewRepo(db *sql.DB, cache *cache.Cache) *Repo {
	return &Repo{db, cache}
}

func (r *Repo) Update(uid string, data string) {
	log := log.Default()
	q, err := r.db.Query(`INSERT INTO public."Order"(
		uid, data)
		VALUES ($1, $2 );`, uid, data)
	if err != nil {
		log.Print(err)
		return
	}
	defer q.Close()
	js := domain.Order{}
	json.Unmarshal([]byte(data), &js)
	r.cache.Update(uid, js)

}

func (r *Repo) GetAll() {
	log := log.Default()
	q, err := r.db.Query(`SELECT * FROM public."Order" `)
	if err != nil {
		log.Print(err)
	}
	uid := ""
	res := ""

	for q.Next() {
		q.Scan(&uid, &res)
		js := domain.Order{}
		json.Unmarshal([]byte(res), &js)
		r.cache.Update(uid, js)

	}

}
