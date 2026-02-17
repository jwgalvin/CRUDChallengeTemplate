package validation_test

import (
	"strings"
	"testing"

	"example.com/crudapp/internal/model"
	"example.com/crudapp/internal/validation"
)

func TestNormalizeItemInput(t *testing.T) {
	tests := []struct {
		name    string
		input   model.ItemInput
		wantErr string
		wantOut model.ItemInput
	}{
		{
			name:    "valid input",
			input:   model.ItemInput{Name: "alpha", Tags: []string{"one", "two"}},
			wantOut: model.ItemInput{Name: "alpha", Tags: []string{"one", "two"}},
		},
		{
			name:    "name is trimmed",
			input:   model.ItemInput{Name: "  alpha  ", Tags: []string{}},
			wantOut: model.ItemInput{Name: "alpha", Tags: []string{}},
		},
		{
			name:    "tags are trimmed",
			input:   model.ItemInput{Name: "alpha", Tags: []string{"  go  ", "  lang  "}},
			wantOut: model.ItemInput{Name: "alpha", Tags: []string{"go", "lang"}},
		},
		{
			name:    "nil tags allowed",
			input:   model.ItemInput{Name: "alpha", Tags: nil},
			wantOut: model.ItemInput{Name: "alpha", Tags: []string{}},
		},
		{
			name:    "empty name",
			input:   model.ItemInput{Name: "", Tags: nil},
			wantErr: "name is required",
		},
		{
			name:    "whitespace-only name",
			input:   model.ItemInput{Name: "   ", Tags: nil},
			wantErr: "name is required",
		},
		{
			name:    "empty string tag",
			input:   model.ItemInput{Name: "alpha", Tags: []string{"valid", ""}},
			wantErr: "tags cannot be empty",
		},
		{
			name:    "whitespace-only tag",
			input:   model.ItemInput{Name: "alpha", Tags: []string{"  "}},
			wantErr: "tags cannot be empty",
		},
		{
			name:    "too many tags",
			input:   model.ItemInput{Name: "alpha", Tags: strings.Fields(strings.Repeat("tag ", 21))},
			wantErr: "tags limit exceeded",
		},
		{
			name:    "exactly 20 tags allowed",
			input:   model.ItemInput{Name: "alpha", Tags: strings.Fields(strings.Repeat("tag ", 20))},
			wantOut: model.ItemInput{Name: "alpha", Tags: strings.Fields(strings.Repeat("tag ", 20))},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			out, err := validation.NormalizeItemInput(tc.input)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Fatalf("expected error %q, got %q", tc.wantErr, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if out.Name != tc.wantOut.Name {
				t.Errorf("name: got %q, want %q", out.Name, tc.wantOut.Name)
			}
			if len(out.Tags) != len(tc.wantOut.Tags) {
				t.Errorf("tags length: got %d, want %d", len(out.Tags), len(tc.wantOut.Tags))
			}
		})
	}
}

func TestParseListFilter(t *testing.T) {
	tests := []struct {
		name       string
		query      string
		wantErr    string
		wantLimit  int
		wantOffset int
		wantName   string
		wantTag    string
	}{
		{
			name:       "defaults applied",
			query:      "",
			wantLimit:  50,
			wantOffset: 0,
		},
		{
			name:       "custom limit",
			query:      "limit=10",
			wantLimit:  10,
			wantOffset: 0,
		},
		{
			name:       "limit capped at max",
			query:      "limit=999",
			wantLimit:  200,
			wantOffset: 0,
		},
		{
			name:       "custom offset",
			query:      "offset=5",
			wantLimit:  50,
			wantOffset: 5,
		},
		{
			name:    "limit zero is invalid",
			query:   "limit=0",
			wantErr: "invalid limit",
		},
		{
			name:    "limit negative is invalid",
			query:   "limit=-1",
			wantErr: "invalid limit",
		},
		{
			name:    "limit non-numeric is invalid",
			query:   "limit=abc",
			wantErr: "invalid limit",
		},
		{
			name:    "offset negative is invalid",
			query:   "offset=-1",
			wantErr: "invalid offset",
		},
		{
			name:    "offset non-numeric is invalid",
			query:   "offset=abc",
			wantErr: "invalid offset",
		},
		{
			name:      "name filter trimmed",
			query:     "name=alpha",
			wantLimit: 50,
			wantName:  "alpha",
		},
		{
			name:      "tag filter",
			query:     "tag=go",
			wantLimit: 50,
			wantTag:   "go",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			values := make(map[string][]string)
			for _, part := range strings.Split(tc.query, "&") {
				if part == "" {
					continue
				}
				kv := strings.SplitN(part, "=", 2)
				if len(kv) == 2 {
					values[kv[0]] = append(values[kv[0]], kv[1])
				}
			}

			filter, err := validation.ParseListFilter(values)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tc.wantErr)
				}
				if err.Error() != tc.wantErr {
					t.Fatalf("expected error %q, got %q", tc.wantErr, err.Error())
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if filter.Limit != tc.wantLimit {
				t.Errorf("limit: got %d, want %d", filter.Limit, tc.wantLimit)
			}
			if filter.Offset != tc.wantOffset {
				t.Errorf("offset: got %d, want %d", filter.Offset, tc.wantOffset)
			}
			if tc.wantName != "" && filter.Name != tc.wantName {
				t.Errorf("name: got %q, want %q", filter.Name, tc.wantName)
			}
			if tc.wantTag != "" && filter.Tag != tc.wantTag {
				t.Errorf("tag: got %q, want %q", filter.Tag, tc.wantTag)
			}
		})
	}
}
