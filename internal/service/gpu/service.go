package gpu

import (
	"context"
	"cquest/internal/models"
	"cquest/internal/repo/gpu"
)

type IGPUService interface {
	GetAllGPUs(ctx context.Context) ([]models.GPU, error)
	AddGPU(ctx context.Context, gpu *models.GPU) error
	UpdateGPU(ctx context.Context, gpu *models.GPU) error
	DeleteGPUByID(ctx context.Context, id int) error
	GetGPUByID(ctx context.Context, id int) (models.GPU, error)
}

type GPUService struct {
	repo gpu.IGPURepo
}

func NewGPUService(repo gpu.IGPURepo) *GPUService {
	return &GPUService{repo: repo}
}
