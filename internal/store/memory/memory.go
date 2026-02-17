package memory

import (
	"context"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store"
)

type Store struct {
	// TODO: Add fields for thread-safe in-memory storage:
	// - map for items (map[string]model.Item)
	// - sync.RWMutex for thread safety
	// - int64 atomic counter for ID generation
}

func New() *Store {
	// TODO: Implement store constructor:
	// 1. Create and return Store with initialized items map
	panic("not implemented")
}

func (s *Store) Create(ctx context.Context, item model.Item) (model.Item, error) {
	// TODO: Implement Create:
	// 1. Generate unique ID (atomic increment, format as "item-{number}")
	// 2. Set CreatedAt and UpdatedAt to time.Now().UTC()
	// 3. Lock for writing and store in map
	// 4. Return created item with populated ID and timestamps
	panic("not implemented")
}

func (s *Store) Get(ctx context.Context, id string) (model.Item, error) {
	// TODO: Implement Get:
	// 1. Read-lock the store
	// 2. Look up item by ID in map
	// 3. Return store.ErrNotFound if not found
	// 4. Otherwise return item and nil error
	panic("not implemented")
}

func (s *Store) List(ctx context.Context, filter store.ListFilter) ([]model.Item, error) {
	// TODO: Implement List with filtering and pagination:
	// 1. Read-lock the store
	// 2. Iterate all items and filter by:
	//    - Name: substring match, case-insensitive
	//    - Tag: exact match, case-insensitive (check if tag in item.Tags)
	// 3. Build results slice
	// 4. Apply pagination:
	//    - Calculate start = filter.Offset
	//    - Calculate end = start + filter.Limit
	//    - Handle bounds (if start > length, return empty)
	// 5. Return paginated slice and nil error
	panic("not implemented")
}

func (s *Store) Update(ctx context.Context, id string, input model.ItemInput) (model.Item, error) {
	// TODO: Implement Update:
	// 1. Write-lock the store
	// 2. Look up item by ID
	// 3. Return store.ErrNotFound if not found
	// 4. Update Name and Tags from input
	// 5. Update UpdatedAt to time.Now().UTC()
	// 6. Keep ID and CreatedAt unchanged
	// 7. Store updated item back in map
	// 8. Return updated item and nil error
	panic("not implemented")
}

func (s *Store) Delete(ctx context.Context, id string) error {
	// TODO: Implement Delete:
	// 1. Write-lock the store
	// 2. Check if item exists with the ID
	// 3. Return store.ErrNotFound if not found
	// 4. Delete from map
	// 5. Return nil error
	panic("not implemented")
}

