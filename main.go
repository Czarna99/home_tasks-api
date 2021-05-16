package main

import (
	"github.com/Pawelek242/home_tasks-api/app"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {
	app.StartApplication()

}
