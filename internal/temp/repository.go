package temp

import (
	"context"
)

type Repository interface {
	GetTest(ctx context.Context) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetTest(ctx context.Context) error {
	return nil
}
