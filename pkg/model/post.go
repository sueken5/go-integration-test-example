package model

import (
	"context"
	"time"
)

type Post struct {
	ID        string
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Repository interface {
	Put(ctx context.Context, src *Post) error
}
