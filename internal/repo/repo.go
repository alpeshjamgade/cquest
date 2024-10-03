package repo

import (
	"context"
	"cquest/internal/clients/db"
	"cquest/internal/models"
)

type IRepo interface {
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

	// laptop
	GetAllLaptops(ctx context.Context) ([]models.Laptop, error)
	AddLaptop(ctx context.Context, gpu *models.Laptop) error
	UpdateLaptop(ctx context.Context, gpu *models.Laptop) error
	DeleteLaptopByID(ctx context.Context, id int) error
	GetLaptopByID(ctx context.Context, id int) (models.Laptop, error)
	GetLaptopsByCPUID(ctx context.Context, cpu_id int) ([]models.Laptop, error)
	GetLaptopsByGPUID(ctx context.Context, gpu_id int) ([]models.Laptop, error)
}

type Repo struct {
	db db.DB
}

func NewRepo(db db.DB) *Repo { return &Repo{db: db} }
