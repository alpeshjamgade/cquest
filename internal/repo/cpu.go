package repo

import (
	"context"
	"cquest/internal/logger"
	"cquest/internal/models"
)

func (repo *Repo) GetAllCPUs(ctx context.Context) ([]models.CPU, error) {
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

func (repo *Repo) AddCPU(ctx context.Context, cpu *models.CPU) error {
	_, err := repo.db.DB().Exec(
		`INSERT INTO cpu(model, cores, threads, cache, base_clock, max_clock, rank, release_date) VALUES($1, $2, $3, $4, $5, $6, $7, $8)`,
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

func (repo *Repo) UpdateCPU(ctx context.Context, cpu *models.CPU) error {
	_, err := repo.db.DB().Exec(
		`UPDATE cpu SET model=$1, cores=$2, threads=$3, cache=$4, base_clock=$5, max_clock=$6, rank=$7, release_date=$8 WHERE ID = $9`,
		cpu.Model,
		cpu.Cores,
		cpu.Threads,
		cpu.Cache,
		cpu.BaseClock,
		cpu.MaxClock,
		cpu.Rank,
		cpu.ReleaseDate,
		cpu.ID)

	if err != nil {
		return err
	}
	return nil
}

func (repo *Repo) DeleteCPUById(ctx context.Context, id int) error {
	_, err := repo.db.DB().Exec(`DELETE FROM cpu WHERE ID = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Repo) GetCPUById(ctx context.Context, id int) (models.CPU, error) {
	var cpu models.CPU
	sqlRow, err := repo.db.DB().Queryx(`SELECT * FROM cpu WHERE ID = $1`, id)
	if err != nil {
		return cpu, err
	}
	err = sqlRow.StructScan(&cpu)
	if err != nil {
		return cpu, err
	}
	return cpu, nil
}
