package sqlite

import (
	"context"
	"crypto/rand"
	"database/sql"
	"errors"
	"fmt"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store"
)

const (
	migrateQuery = `CREATE TABLE IF NOT EXISTS items (
        id         TEXT PRIMARY KEY,
        name       TEXT NOT NULL,
        tags       TEXT NOT NULL DEFAULT '[]',
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    )`
	insertQuery         = `INSERT INTO items (id, name, tags, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`
	getQuery            = `SELECT id, name, tags, created_at, updated_at FROM items WHERE id = ?`
	listBaseQuery       = `SELECT id, name, tags, created_at, updated_at FROM items WHERE 1=1`
	listTagsFilterQuery = ` AND EXISTS (SELECT 1 FROM json_each(tags) WHERE LOWER(value) = LOWER(?))`
	listOrderbyQuery    = ` ORDER BY created_at ASC LIMIT ? OFFSET ?`
)

type Store struct {
	myDb *sql.DB
}

func New(dsn string) *Store {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		panic(fmt.Sprintf("sqlite open: %v", err))
	}
	if err := migrate(db); err != nil {
		panic(fmt.Sprintf("sqlite migrate: %v", err))
	}
	return &Store{myDb: db}
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS items (
        id         TEXT PRIMARY KEY,
        name       TEXT NOT NULL,
        tags       TEXT NOT NULL DEFAULT '[]',
        created_at DATETIME NOT NULL,
        updated_at DATETIME NOT NULL
    )`)
	return err
}

func generateID() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return fmt.Sprintf("item-%x", b), nil
}

func (s *Store) Create(_ context.Context, _ model.Item) (model.Item, error) {
	// generate UUID
	// create tags (should be empty if no tags provided and will need meet a not null constraint)
	// marshal tags to JSON
	// get current timestamp
	// set the ID, tags, and timestamps on the item
	return model.Item{}, errors.New("sqlite store not implemented")
}

func (s *Store) Get(_ context.Context, _ string) (model.Item, error) {
	//
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
