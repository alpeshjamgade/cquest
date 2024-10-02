package gpu

import (
	"context"
	"cquest/internal/clients/db"
	"cquest/internal/models"
)

type IGPURepo interface {
	GetAllGPUs(ctx context.Context) ([]models.GPU, error)
	AddGPU(ctx context.Context, gpu *models.GPU) error
	UpdateGPU(ctx context.Context, gpu *models.GPU) error
	DeleteGPUByID(ctx context.Context, id int) error
	GetGPUByID(ctx context.Context, id int) (models.GPU, error)
}

type GPURepo struct {
	db db.DB
}

func NewGPURepo(db db.DB) *GPURepo {
	return &GPURepo{db: db}
}
