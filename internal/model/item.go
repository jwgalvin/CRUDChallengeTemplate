package model

import "time"

type Item struct {
    ID        string    `json:"id"`
    Name      string    `json:"name"`
    Tags      []string  `json:"tags"`
    CreatedAt time.Time `json:"createdAt"`
    UpdatedAt time.Time `json:"updatedAt"`
}

type ItemInput struct {
    Name string   `json:"name"`
    Tags []string `json:"tags"`
}
