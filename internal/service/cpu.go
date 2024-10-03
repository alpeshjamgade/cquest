package service

import (
	"context"
	"cquest/internal/models"
)

func (svc *Service) GetAllCPUs(ctx context.Context) ([]models.CPU, error) {
	cpus, err := svc.repo.GetAllCPUs(ctx)

	if err != nil {
		return nil, err
	}

	return cpus, nil
}

func (svc *Service) AddCPU(ctx context.Context, cpu *models.CPU) error {
	err := svc.repo.AddCPU(ctx, cpu)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateCPU(ctx context.Context, cpu *models.CPU) error {
	err := svc.repo.UpdateCPU(ctx, cpu)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeleteCPUById(ctx context.Context, id int) error {
	err := svc.repo.DeleteCPUById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) GetCPUById(ctx context.Context, id int) (models.CPU, error) {
	var cpu models.CPU
	cpu, err := svc.repo.GetCPUById(ctx, id)
	if err != nil {
		return cpu, err
	}
	return cpu, nil
}
