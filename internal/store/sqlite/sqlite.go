package sqlite

import (
	"context"
	"errors"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store"
)

type Store struct {
    dsn string
}

func New(dsn string) *Store {
    return &Store{dsn: dsn}
}

func (s *Store) Create(_ context.Context, _ model.Item) (model.Item, error) {
    return model.Item{}, errors.New("sqlite store not implemented")
}

func (s *Store) Get(_ context.Context, _ string) (model.Item, error) {
    return model.Item{}, errors.New("sqlite store not implemented")
}

func (s *Store) List(_ context.Context, _ store.ListFilter) ([]model.Item, error) {
    return nil, errors.New("sqlite store not implemented")
}

func (s *Store) Update(_ context.Context, _ string, _ model.ItemInput) (model.Item, error) {
    return model.Item{}, errors.New("sqlite store not implemented")
}

func (s *Store) Delete(_ context.Context, _ string) error {
    return errors.New("sqlite store not implemented")
}
