package repo

import (
	"api/interal/pkg/cache"
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

func (r Repo) Get(uid string) string {
	log := log.Default()
	//r.cache.Get(uid)
	/*q, err := r.db.Query(`SELECT data FROM public."Orders" WHERE uid = $1`, uid)
	if err != nil {
		return "nothing"
	}
	defer q.Close()
	res := ""
	for q.Next() {
		err = q.Scan(&res)
		if err != nil {
			log.Print("error get data")
		}
	}*/
	res, err := json.Marshal(r.cache.Get(uid))
	if err != nil {
		log.Print(err)
		return ""
	}
	return string(res)
}
