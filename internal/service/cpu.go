package service

import (
	"context"
	"cquest/internal/models"
)

func (svc *CPUService) GetAllCPUs(ctx context.Context) ([]models.CPU, error) {
	cpus, err := svc.repo.GetAllCPUs(ctx)

	if err != nil {
		return nil, err
	}

	return cpus, nil
}

func (svc *CPUService) AddCPU(ctx context.Context, cpu *models.CPU) error {
	err := svc.repo.AddCPU(ctx, cpu)
	if err != nil {
		return err
	}

	return nil
}

func (svc *CPUService) UpdateCPU(ctx context.Context, cpu *models.CPU) error {
	err := svc.repo.UpdateCPU(ctx, cpu)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CPUService) DeleteCPUByID(ctx context.Context, id int) error {
	err := svc.repo.DeleteCPUByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *CPUService) GetCPUByID(ctx context.Context, id int) (models.CPU, error) {
	var cpu models.CPU
	cpu, err := svc.repo.GetCPUByID(ctx, id)
	if err != nil {
		return cpu, err
	}
	return cpu, nil
}
