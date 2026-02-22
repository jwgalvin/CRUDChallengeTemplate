package store

import (
	"context"
	"errors"

	"example.com/crudapp/internal/model"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrUnsupportedStore = errors.New("unsupported store")
)

type ListFilter struct {
	Name   string
	Tag    string
	Limit  int
	Offset int
}

type Store interface {
	Create(ctx context.Context, item model.Item) (model.Item, error)
	Get(ctx context.Context, id string) (model.Item, error)
	List(ctx context.Context, filter ListFilter) ([]model.Item, error)
}
