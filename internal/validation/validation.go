package validation

import (
	"net/url"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store"
)

const (
	defaultLimit = 50
	maxLimit     = 200
)

func NormalizeItemInput(input model.ItemInput) (model.ItemInput, error) {
	// TODO: Implement input normalization and validation:
	// 1. Trim whitespace from name
	// 2. Return error "name is required" if name is empty after trimming
	// 3. Return error "tags limit exceeded" if len(tags) > 20
	// 4. For each tag:
	//    a. Trim whitespace
	//    b. Return error "tags cannot be empty" if trimmed tag is empty
	//    c. Otherwise add to normalized list
	// 5. Return normalized ItemInput with trimmed name and tags
	panic("not implemented")
}

func ParseListFilter(values url.Values) (store.ListFilter, error) {
	// TODO: Implement query parameter parsing:
	// 1. Create ListFilter with defaults:
	//    - Name: trim and get "name" param (default "")
	//    - Tag: trim and get "tag" param (default "")
	//    - Limit: defaultLimit (50)
	//    - Offset: 0
	// 2. If "limit" param exists:
	//    a. Parse as integer
	//    b. Return error "invalid limit" if not parseable or < 1
	//    c. Cap at maxLimit (200)
	// 3. If "offset" param exists:
	//    a. Parse as integer
	//    b. Return error "invalid offset" if not parseable or < 0
	// 4. Return completed filter and nil error
	panic("not implemented")
}
