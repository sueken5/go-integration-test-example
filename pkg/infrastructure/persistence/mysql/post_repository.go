package mysql

import (
	"context"
	"fmt"

	"github.com/sueken5/go-integration-test-example/pkg/model"
	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) Put(ctx context.Context, src *model.Post) error {
	if err := r.db.Create(src).Error; err != nil {
		return fmt.Errorf("post repository put err: %w", err)
	}

	return nil
}

func NewPostRepository(
	db *gorm.DB,
) model.Repository {
	return &PostRepository{db: db}
}
