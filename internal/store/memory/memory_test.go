package memory

import (
	"context"
	"testing"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store"
)

func TestMemoryStoreCRUD(t *testing.T) {
    st := New()

    created, err := st.Create(context.Background(), model.Item{Name: "alpha", Tags: []string{"one"}})
    if err != nil {
        t.Fatalf("create: %v", err)
    }

    fetched, err := st.Get(context.Background(), created.ID)
    if err != nil {
        t.Fatalf("get: %v", err)
    }

    if fetched.Name != "alpha" {
        t.Fatalf("expected name alpha, got %s", fetched.Name)
    }

    updated, err := st.Update(context.Background(), created.ID, model.ItemInput{Name: "beta", Tags: []string{"two"}})
    if err != nil {
        t.Fatalf("update: %v", err)
    }

    if updated.Name != "beta" {
        t.Fatalf("expected name beta, got %s", updated.Name)
    }

    if err := st.Delete(context.Background(), created.ID); err != nil {
        t.Fatalf("delete: %v", err)
    }

    if _, err := st.Get(context.Background(), created.ID); err == nil {
        t.Fatalf("expected not found error")
    }
}

func TestMemoryStoreListFilter(t *testing.T) {
	st := New()

	_, _ = st.Create(context.Background(), model.Item{Name: "alpha", Tags: []string{"a"}})
	_, _ = st.Create(context.Background(), model.Item{Name: "beta", Tags: []string{"b"}})

	results, err := st.List(context.Background(), store.ListFilter{Name: "alp", Limit: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}

	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
}

func TestGetNotFound(t *testing.T) {
	st := New()

	_, err := st.Get(context.Background(), "nonexistent-id")
	if err != store.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestUpdateNotFound(t *testing.T) {
	st := New()

	_, err := st.Update(context.Background(), "nonexistent-id", model.ItemInput{Name: "test"})
	if err != store.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestDeleteNotFound(t *testing.T) {
	st := New()

	err := st.Delete(context.Background(), "nonexistent-id")
	if err != store.ErrNotFound {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestListByTag(t *testing.T) {
	st := New()

	st.Create(context.Background(), model.Item{Name: "go-app", Tags: []string{"golang", "web"}})
	st.Create(context.Background(), model.Item{Name: "python-api", Tags: []string{"python", "web"}})
	st.Create(context.Background(), model.Item{Name: "rust-cli", Tags: []string{"rust"}})

	results, err := st.List(context.Background(), store.ListFilter{Tag: "golang", Limit: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result with tag golang, got %d", len(results))
	}

	results, err = st.List(context.Background(), store.ListFilter{Tag: "web", Limit: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 2 {
		t.Fatalf("expected 2 results with tag web, got %d", len(results))
	}

	results, err = st.List(context.Background(), store.ListFilter{Tag: "nonexistent", Limit: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected 0 results with tag nonexistent, got %d", len(results))
	}
}

func TestListPagination(t *testing.T) {
	st := New()

	for i := 0; i < 10; i++ {
		st.Create(context.Background(), model.Item{Name: "item-" + string(rune(i)), Tags: []string{}})
	}

	// Test limit
	results, err := st.List(context.Background(), store.ListFilter{Limit: 5})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 5 {
		t.Fatalf("expected 5 results with limit 5, got %d", len(results))
	}

	// Test offset
	results, err = st.List(context.Background(), store.ListFilter{Limit: 5, Offset: 5})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 5 {
		t.Fatalf("expected 5 results with offset 5, got %d", len(results))
	}

	// Test offset beyond results
	results, err = st.List(context.Background(), store.ListFilter{Limit: 10, Offset: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected 0 results with offset 10 on 10 items, got %d", len(results))
	}
}

func TestListCaseInsensitive(t *testing.T) {
	st := New()

	st.Create(context.Background(), model.Item{Name: "GoLang-API", Tags: []string{"Backend"}})
	st.Create(context.Background(), model.Item{Name: "Python-App", Tags: []string{"Frontend"}})

	// Test name case insensitivity
	results, err := st.List(context.Background(), store.ListFilter{Name: "golang", Limit: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result for case-insensitive name search, got %d", len(results))
	}

	// Test tag case insensitivity
	results, err = st.List(context.Background(), store.ListFilter{Tag: "backend", Limit: 10})
	if err != nil {
		t.Fatalf("list: %v", err)
	}
	if len(results) != 1 {
		t.Fatalf("expected 1 result for case-insensitive tag search, got %d", len(results))
	}
}

func TestUpdatePreservesCreatedAt(t *testing.T) {
	st := New()

	created, _ := st.Create(context.Background(), model.Item{Name: "original", Tags: []string{"a"}})
	originalCreatedAt := created.CreatedAt

	updated, _ := st.Update(context.Background(), created.ID, model.ItemInput{Name: "modified", Tags: []string{"b"}})

	if !updated.CreatedAt.Equal(originalCreatedAt) {
		t.Fatalf("update should not change createdAt")
	}
	if updated.UpdatedAt.Before(originalCreatedAt) {
		t.Fatalf("updatedAt should be after createdAt")
	}
}

func TestCreateGeneratesUniqueIDs(t *testing.T) {
	st := New()

	ids := make(map[string]bool)
	for i := 0; i < 10; i++ {
		item, _ := st.Create(context.Background(), model.Item{Name: "item", Tags: []string{}})
		if ids[item.ID] {
			t.Fatalf("duplicate ID generated: %s", item.ID)
		}
		ids[item.ID] = true
	}
}

func TestListEmptyStore(t *testing.T) {
	st := New()

	results, err := st.List(context.Background(), store.ListFilter{Limit: 10})
	if err != nil {
		t.Fatalf("list on empty store: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected 0 results on empty store, got %d", len(results))
	}
}
