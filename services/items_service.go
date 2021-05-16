package services

import (
	"github.com/Pawelek242/home_oauth-api/src/utils/errors"
	"github.com/Pawelek242/home_tasks-api/items"
)

var (
	err          []string
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *errors.RestErr)
	Get(string) (*items.Item, *errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, errors.NewBadRequest(append(err, "not implemented"))
}

func (s *itemsService) Get(string) (*items.Item, *errors.RestErr) {
	return nil, errors.NewBadRequest(append(err, "not implemented"))
}
