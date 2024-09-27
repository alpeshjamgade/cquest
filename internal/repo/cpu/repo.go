package cpu

import (
	"context"
	"cquest/internal/models"
)

type Repo interface {
	GetAllCPUs(ctx context.Context) ([]models.CPU, error)
	AddCPU(ctx context.Context, cpu models.CPU) error
	UpdateCPU(ctx context.Context, cpu models.CPU) error
	DeleteCPU(ctx context.Context, cpu models.CPU) error
	GetCPU(ctx context.Context, cpu models.CPU) (models.CPU, error)
}

type CPURepo struct {
}
