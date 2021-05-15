package controllers

import (
	"net/http"

	"github.com/Pawelek242/home_oauth-go/oauth"
)

func Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		return
	}
}

func Get(w http.ResponseWriter, r *http.Request) {

}
