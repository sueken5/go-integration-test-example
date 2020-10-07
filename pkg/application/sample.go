package application

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sueken5/go-integration-test-example/pkg/apis/sample"
	"github.com/sueken5/go-integration-test-example/pkg/model"
)

type Sample struct {
	repo model.Repository
}

func (s *Sample) Put(ctx context.Context, request *sample.PutRequest) (*sample.PutResponse, error) {
	post := &model.Post{
		ID:        uuid.NewV4().String(),
		Message:   request.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Put(ctx, post); err != nil {
		return nil, fmt.Errorf("sample put err: %w", err)
	}

	return &sample.PutResponse{}, nil
}

func NewSample(
	repo model.Repository,
) sample.SampleServer {
	return &Sample{repo: repo}
}
