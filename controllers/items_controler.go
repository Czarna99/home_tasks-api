package controllers

import (
	"fmt"
	"net/http"

	"github.com/Pawelek242/home_oauth-go/oauth"
	"github.com/Pawelek242/home_tasks-api/items"
	"github.com/Pawelek242/home_tasks-api/services"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct{}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO return error json to the user
		return
	}

	item := items.Item{
		Author: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		//TODO return error json to the user
	}

	fmt.Println(result)
	//Return created item as json with HTTP status 201 - created

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
