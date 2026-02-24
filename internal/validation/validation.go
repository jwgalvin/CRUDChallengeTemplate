package validation

import (
	"fmt"
	"net/url"
	"strconv"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/store"
)

const (
	defaultLimit = 50
	maxLimit     = 200
	tagLimit     = 20
)

func NormalizeItemInput(input model.ItemInput) (model.ItemInput, error) {
	panic("not implemented")
}

func ParseListFilter(values url.Values) (store.ListFilter, error) {
	panic("not implemented")
}

func parseLimit(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < 1 {
		return 0, fmt.Errorf("invalid limit")
	}
	if n > maxLimit {
		n = maxLimit
	}
	return n, nil
}

func parseOffset(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil || n < 0 {
		return 0, fmt.Errorf("invalid offset")
	}
	return n, nil
}
