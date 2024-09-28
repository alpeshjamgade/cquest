package cpu

import (
	"context"
	"cquest/internal/logger"
	"cquest/internal/models"
)

func (repo *CPURepo) GetAllCPUs(ctx context.Context) ([]models.CPU, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)
	var cpus []models.CPU
	sqlRow, err := repo.db.DB().Queryx(`SELECT * FROM cpu ORDER BY id DESC`)
	if err != nil {
		Logger.Errorw("error while fetching all cpus", "error", err)
		return nil, err
	}

	for sqlRow.Next() {
		var cpu models.CPU
		err = sqlRow.StructScan(&cpu)
		if err != nil {
			Logger.Errorw("error while scanning the row", "error", err)
			return nil, err
		}
		cpus = append(cpus, cpu)
	}
	return cpus, nil
}

func (repo *CPURepo) AddCPU(ctx context.Context, cpu models.CPU) error {
	_, err := repo.db.DB().Exec(
		`INSERT INTO cpu(Model, Cores, Threads, Cache, BaseClock, MaxClock, Rank, ReleaseDate) VALUES($1, $2, $3, $4, $5, $6, $7)`,
		cpu.Model,
		cpu.Cores,
		cpu.Threads,
		cpu.Cache,
		cpu.BaseClock,
		cpu.MaxClock,
		cpu.Rank,
		cpu.ReleaseDate)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CPURepo) UpdateCPU(ctx context.Context, cpu models.CPU) error {
	_, err := repo.db.DB().Exec(
		`UPDATE cpu SET model=$1, cores=$2, thread=$3, cache=$3, base_clock=$4, max_clock=$5, rank=$6, release_date=$7 WHERE ID = $8`,
		cpu.Model,
		cpu.Cores,
		cpu.Threads,
		cpu.Threads,
		cpu.Cache,
		cpu.MaxClock,
		cpu.Rank,
		cpu.ReleaseDate,
		cpu.ID)

	if err != nil {
		return err
	}
	return nil
}

func (repo *CPURepo) DeleteCPUByID(ctx context.Context, id int) error {
	_, err := repo.db.DB().Exec(`DELETE FROM cpu WHERE ID = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CPURepo) GetCPUByID(ctx context.Context, id int) (models.CPU, error) {
	var cpu models.CPU
	sqlRow, err := repo.db.DB().Queryx(`SELECT * FROM cpu WHERE ID = $1`, id)
	if err != nil {
		return cpu, err
	}
	for sqlRow.Next() {
		err = sqlRow.StructScan(&cpu)
		if err != nil {
			return cpu, err
		}
	}
	return cpu, nil
}
