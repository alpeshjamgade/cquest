package repo

import (
	"context"
	"cquest/internal/logger"
	"cquest/internal/models"
)

func (repo *GPURepo) GetAllGPUs(ctx context.Context) ([]models.GPU, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)

	var gpus []models.GPU
	sqlRow, err := repo.db.DB().Queryx(`SELECT * FROM gpu ORDER BY id ASC`)
	if err != nil {
		Logger.Errorf("Error fetching gpus, %s", err)
		return gpus, err
	}

	for sqlRow.Next() {
		var gpu models.GPU
		err := sqlRow.StructScan(&gpu)
		if err != nil {
			Logger.Errorf("Error scanning fetched GPU records, %s", err)
			return gpus, err
		}
		gpus = append(gpus, gpu)
	}

	return gpus, nil
}

func (repo *GPURepo) AddGPU(ctx context.Context, gpu *models.GPU) error {
	Logger := logger.CreateLoggerWithCtx(ctx)
	_, err := repo.db.DB().Exec(
		`INSERT INTO gpu(model, memory, clock_speed, memory_type, tdp, architecture, rank, release_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
		gpu.Model,
		gpu.Memory,
		gpu.ClockSpeed,
		gpu.MemoryType,
		gpu.TDP,
		gpu.Architecture,
		gpu.Rank,
		gpu.ReleaseDate)

	if err != nil {
		Logger.Errorf("Error inserting GPU, %s", err)
		return err
	}
	return nil
}

func (repo *GPURepo) UpdateGPU(ctx context.Context, gpu *models.GPU) error {
	Logger := logger.CreateLoggerWithCtx(ctx)
	_, err := repo.db.DB().Exec(
		`UPDATE gpu SET model=$1, memory=$2, clock_speed=$3, memory_type=$4, tdp=$5, architecture=$6, rank=$7, release_date=$8 WHERE id=$9`,
		gpu.Model, gpu.Memory, gpu.ClockSpeed, gpu.MemoryType, gpu.TDP, gpu.Architecture, gpu.Rank, gpu.ReleaseDate, gpu.ID,
	)

	if err != nil {
		Logger.Errorf("Error updating gpu, %s", err)
		return err
	}

	return nil
}

func (repo *GPURepo) DeleteGPUByID(ctx context.Context, id int) error {
	Logger := logger.CreateLoggerWithCtx(ctx)
	_, err := repo.db.DB().Exec(
		`DELETE FROM gpu WHERE id=$1`,
		id,
	)

	if err != nil {
		Logger.Errorf("Error deleting GPU, %s", err)
		return err
	}

	return nil
}

func (repo *GPURepo) GetGPUByID(ctx context.Context, id int) (models.GPU, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)
	var gpu models.GPU
	sqlRow := repo.db.DB().QueryRowx(
		`SELECT * FROM gpu WHERE id=$1`,
		id,
	)

	err := sqlRow.StructScan(&gpu)
	if err != nil {
		Logger.Errorf("Error scanning fetched GPU record, %s", err)
		return gpu, err
	}

	return gpu, err
}
