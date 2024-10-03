package service

import (
	"context"
	"cquest/internal/models"
)

func (svc *Service) GetAllGPUs(ctx context.Context) ([]models.GPU, error) {
	gpus, err := svc.repo.GetAllGPUs(ctx)
	if err != nil {
		return gpus, err
	}

	return gpus, nil
}

func (svc *Service) AddGPU(ctx context.Context, gpu *models.GPU) error {
	err := svc.repo.AddGPU(ctx, gpu)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateGPU(ctx context.Context, gpu *models.GPU) error {
	err := svc.repo.UpdateGPU(ctx, gpu)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteGPUById(ctx context.Context, id int) error {
	err := svc.repo.DeleteGPUById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) GetGPUById(ctx context.Context, id int) (models.GPU, error) {
	gpu, err := svc.repo.GetGPUById(ctx, id)
	if err != nil {
		return gpu, err
	}

	return gpu, nil
}
