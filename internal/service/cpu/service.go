package cpu

import (
	"context"
	"cquest/internal/models"
	"cquest/internal/repo/cpu"
)

type ICPUService interface {
	GetAllCPUs(ctx context.Context) ([]models.CPU, error)
	AddCPU(ctx context.Context, cpu *models.CPU) error
	UpdateCPU(ctx context.Context, cpu *models.CPU) error
	DeleteCPUByID(ctx context.Context, id int) error
	GetCPUByID(ctx context.Context, id int) (models.CPU, error)
}

type CPUService struct {
	repo cpu.ICPURepo
}

func NewCPUService(repo cpu.ICPURepo) *CPUService {
	service := &CPUService{
		repo: repo,
	}

	return service
}
