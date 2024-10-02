package gpu

import (
	"context"
	"cquest/internal/models"
)

func (svc GPUService) GetAllGPUs(ctx context.Context) ([]models.GPU, error) {
	gpus, err := svc.repo.GetAllGPUs(ctx)
	if err != nil {
		return gpus, err
	}

	return gpus, nil
}

func (svc GPUService) AddGPU(ctx context.Context, gpu *models.GPU) error {
	err := svc.repo.AddGPU(ctx, gpu)
	if err != nil {
		return err
	}

	return nil
}

func (svc GPUService) UpdateGPU(ctx context.Context, gpu *models.GPU) error {
	err := svc.repo.UpdateGPU(ctx, gpu)
	if err != nil {
		return err
	}

	return nil
}

func (svc GPUService) DeleteGPUByID(ctx context.Context, id int) error {
	err := svc.repo.DeleteGPUByID(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (svc GPUService) GetGPUByID(ctx context.Context, id int) (models.GPU, error) {
	gpu, err := svc.repo.GetGPUByID(ctx, id)
	if err != nil {
		return gpu, err
	}

	return gpu, nil
}
