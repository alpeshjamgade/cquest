package service

import (
	"context"
	"cquest/internal/models"
	"cquest/internal/repo"
)

type IService interface {
	// cpu
	GetAllCPUs(ctx context.Context) ([]models.CPU, error)
	AddCPU(ctx context.Context, cpu *models.CPU) error
	UpdateCPU(ctx context.Context, cpu *models.CPU) error
	DeleteCPUByID(ctx context.Context, id int) error
	GetCPUByID(ctx context.Context, id int) (models.CPU, error)

	// gpu
	GetAllGPUs(ctx context.Context) ([]models.GPU, error)
	AddGPU(ctx context.Context, gpu *models.GPU) error
	UpdateGPU(ctx context.Context, gpu *models.GPU) error
	DeleteGPUByID(ctx context.Context, id int) error
	GetGPUByID(ctx context.Context, id int) (models.GPU, error)
}

type Service struct {
	repo repo.IRepo
}

func NewService(repo repo.IRepo) *Service {
	return &Service{repo: repo}
}
