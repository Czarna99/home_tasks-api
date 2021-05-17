package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Pawelek242/home_oauth-go/oauth"
	"github.com/Pawelek242/home_tasks-api/domain/items"
	"github.com/Pawelek242/home_tasks-api/services"
	"github.com/Pawelek242/home_tasks-api/utils/http_utils"
	"github.com/Pawelek242/home_utils-go/rest_errors"
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
		http_utils.RespondError(w, err)
		return
	}
	var itemRequest items.Item

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json")
		http_utils.RespondError(w, respErr)
		return

	}
	itemRequest.Author = oauth.GetClientId(r)
	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, createErr)
		return

	}

	http_utils.RespondJson(w, http.StatusCreated, result)

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
