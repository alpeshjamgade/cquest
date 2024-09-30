package cpu

import (
	"context"
	"cquest/internal/clients/db"
	"cquest/internal/models"
)

type ICPURepo interface {
	GetAllCPUs(ctx context.Context) ([]models.CPU, error)
	AddCPU(ctx context.Context, cpu *models.CPU) error
	UpdateCPU(ctx context.Context, cpu *models.CPU) error
	DeleteCPUByID(ctx context.Context, id int) error
	GetCPUByID(ctx context.Context, id int) (models.CPU, error)
}

type CPURepo struct {
	db db.DB
}

func NewCPURepo(db db.DB) *CPURepo {
	repo := &CPURepo{db: db}
	return repo
}
