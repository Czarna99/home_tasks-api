package app

import (
	"net/http"
	"time"

	"github.com/Pawelek242/home_tasks-api/clients/elasticsearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticsearch.Init()
	mapURLs()

	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Millisecond,
		IdleTimeout:  60 * time.Millisecond,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
