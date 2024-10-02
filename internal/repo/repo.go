package repo

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

type IGPURepo interface {
	GetAllGPUs(ctx context.Context) ([]models.GPU, error)
	AddGPU(ctx context.Context, gpu *models.GPU) error
	UpdateGPU(ctx context.Context, gpu *models.GPU) error
	DeleteGPUByID(ctx context.Context, id int) error
	GetGPUByID(ctx context.Context, id int) (models.GPU, error)
}

type CPURepo struct {
	db db.DB
}

type GPURepo struct {
	db db.DB
}

func NewCPURepo(db db.DB) *CPURepo { return &CPURepo{db: db} }

func NewGPURepo(db db.DB) *GPURepo { return &GPURepo{db: db} }
