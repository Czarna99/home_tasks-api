package app

import (
	"net/http"

	"github.com/Pawelek242/home_tasks-api/controllers"
)

func mapURLs() {
	router.HandleFunc("/items/", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
}
