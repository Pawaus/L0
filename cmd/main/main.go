package main

import (
	"api/interal/app/get"
	"api/interal/app/update"
	"api/interal/pkg/cache"
	"api/interal/pkg/nats"
	"api/interal/pkg/postgres"
	"log"
	"net/http"
)

func main() {
	logger := log.Default()
	pgCfg := postgres.PgConfig{
		Host:     "192.168.101.1",
		Port:     "5432",
		Username: "postgres",
		Password: "Passw0rd",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	db, err := postgres.ConnDB(pgCfg)

	if err != nil {
		logger.Panic(err)
	}

	ch := cache.New()

	router := get.Setup(db, ch)
	stream := update.Setup(db, ch)
	stream.UpdateBD()
	sub, err := nats.Connect("test-cluster", "client-123")
	(*sub).Subscribe("foo", stream.UpdateChannel)

	http.HandleFunc("/get", router.Get)

	http.ListenAndServe("192.168.101.1:5005", nil)

}
