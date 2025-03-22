package multimodal

import (
	"context"
)

type Service interface {
	GetTest(ctx context.Context) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetTest(ctx context.Context) error {
	return s.repository.GetTest(ctx)
}
