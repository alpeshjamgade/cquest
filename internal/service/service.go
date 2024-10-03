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
	DeleteCPUById(ctx context.Context, id int) error
	GetCPUById(ctx context.Context, id int) (models.CPU, error)

	// gpu
	GetAllGPUs(ctx context.Context) ([]models.GPU, error)
	AddGPU(ctx context.Context, gpu *models.GPU) error
	UpdateGPU(ctx context.Context, gpu *models.GPU) error
	DeleteGPUById(ctx context.Context, id int) error
	GetGPUById(ctx context.Context, id int) (models.GPU, error)

	// laptop
	GetAllLaptops(ctx context.Context) ([]models.Laptop, error)
	AddLaptop(ctx context.Context, gpu *models.Laptop) error
	UpdateLaptop(ctx context.Context, gpu *models.Laptop) error
	DeleteLaptopById(ctx context.Context, id int) error
	GetLaptopById(ctx context.Context, id int) (models.Laptop, error)
	GetLaptopsByCPUId(ctx context.Context, cpu_id int) ([]models.Laptop, error)
	GetLaptopsByGPUId(ctx context.Context, gpu_id int) ([]models.Laptop, error)
}

type Service struct {
	repo repo.IRepo
}

func NewService(repo repo.IRepo) *Service {
	return &Service{repo: repo}
}
