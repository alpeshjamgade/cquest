package service

import (
	"context"
	"cquest/internal/models"
)

func (svc *Service) GetAllLaptops(ctx context.Context) ([]models.Laptop, error) {
	laptops, err := svc.repo.GetAllLaptops(ctx)
	if err != nil {
		return laptops, err
	}

	return laptops, nil
}

func (svc *Service) AddLaptop(ctx context.Context, laptop *models.Laptop) error {
	err := svc.AddLaptop(ctx, laptop)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) UpdateLaptop(ctx context.Context, laptop *models.Laptop) error {
	err := svc.UpdateLaptop(ctx, laptop)
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) DeleteLaptopById(ctx context.Context, id int) error {
	err := svc.DeleteLaptopById(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) GetLaptopById(ctx context.Context, id int) (models.Laptop, error) {
	laptop, err := svc.GetLaptopById(ctx, id)
	if err != nil {
		return laptop, err
	}

	return laptop, nil
}

func (svc *Service) GetLaptopsByCPUId(ctx context.Context, cpu_id int) ([]models.Laptop, error) {
	laptops, err := svc.GetLaptopsByCPUId(ctx, cpu_id)
	if err != nil {
		return laptops, err
	}

	return laptops, nil
}

func (svc *Service) GetLaptopsByGPUId(ctx context.Context, gpu_id int) ([]models.Laptop, error) {
	laptops, err := svc.GetLaptopsByGPUId(ctx, gpu_id)
	if err != nil {
		return laptops, err
	}

	return laptops, nil
}
