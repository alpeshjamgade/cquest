package service

import (
	"context"
	"cquest/internal/models"
	"cquest/internal/repo"
)

type ICPUService interface {
	GetAllCPUs(ctx context.Context) ([]models.CPU, error)
	AddCPU(ctx context.Context, cpu *models.CPU) error
	UpdateCPU(ctx context.Context, cpu *models.CPU) error
	DeleteCPUByID(ctx context.Context, id int) error
	GetCPUByID(ctx context.Context, id int) (models.CPU, error)
}

type IGPUService interface {
	GetAllGPUs(ctx context.Context) ([]models.GPU, error)
	AddGPU(ctx context.Context, gpu *models.GPU) error
	UpdateGPU(ctx context.Context, gpu *models.GPU) error
	DeleteGPUByID(ctx context.Context, id int) error
	GetGPUByID(ctx context.Context, id int) (models.GPU, error)
}

type CPUService struct {
	repo repo.ICPURepo
}

type GPUService struct {
	repo repo.IGPURepo
}

func NewCPUService(repo repo.ICPURepo) *CPUService { return &CPUService{repo: repo} }

func NewGPUService(repo repo.IGPURepo) *GPUService {
	return &GPUService{repo: repo}
}
