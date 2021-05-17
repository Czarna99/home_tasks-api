package services

import (
	"github.com/Pawelek242/home_tasks-api/domain/items"
	"github.com/Pawelek242/home_utils-go/rest_errors"
)

var (
	err          []string
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewNotFoundError("not implemented")
}

func (s *itemsService) Get(string) (*items.Item, rest_errors.RestErr) {
	return nil, rest_errors.NewNotFoundError("not implemented")
}
