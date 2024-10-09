package repo

import (
	"context"
	"cquest/internal/logger"
	"cquest/internal/models"
	"database/sql"
	"errors"
)

func (repo *Repo) GetAllLaptops(ctx context.Context) ([]models.Laptop, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)

	var laptops []models.Laptop
	sqlRow, err := repo.db.DB().Queryx(`SELECT * FROM laptop ORDER BY id ASC`)
	if err != nil {
		Logger.Errorf("Error fetching laptops, %s", err)
		return laptops, err
	}

	for sqlRow.Next() {
		var laptop models.Laptop
		err := sqlRow.StructScan(&laptop)
		if err != nil {
			Logger.Errorf("Error scanning fetched Laptop records, %s", err)
			return laptops, err
		}
		laptops = append(laptops, laptop)
	}

	return laptops, nil
}

func (repo *Repo) AddLaptop(ctx context.Context, laptop *models.Laptop) error {
	Logger := logger.CreateLoggerWithCtx(ctx)
	_, err := repo.db.DB().Exec(
		`INSERT INTO laptop(
                brand, model, cpu_id, gpu_id, ram, ssd, hdd, usb_c, usb_a, hdmi, ethernet,
                headphone_jack, weight, price, rank, release_date)
				VALUES($1,	 $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
		laptop.Brand,
		laptop.Model,
		laptop.CPUId,
		laptop.GPUId,
		laptop.RAM,
		laptop.SSD,
		laptop.HDD,
		laptop.UsbC,
		laptop.UsbA,
		laptop.HDMI,
		laptop.Ethernet,
		laptop.HeadphoneJack,
		laptop.Weight,
		laptop.PriceInRupees,
		laptop.Rank,
		laptop.ReleaseDate)

	if err != nil {
		Logger.Errorf("Error inserting Laptop, %s", err)
		return err
	}
	return nil
}

func (repo *Repo) UpdateLaptop(ctx context.Context, laptop *models.Laptop) error {
	Logger := logger.CreateLoggerWithCtx(ctx)
	_, err := repo.db.DB().Exec(
		`UPDATE laptop SET brand = $1, model = $2, cpu_id = $3, gpu_id = $4, ram = $5, ssd = $6, hdd = $7, usb_c = $8, usb_a = $9, hdmi = $10, ethernet = $11, headphone_jack = $12, weight = $13, price = $14, rank = $15, release_date = $16, WHERE id=$17`,
		laptop.Brand,
		laptop.Model,
		laptop.CPUId,
		laptop.GPUId,
		laptop.RAM,
		laptop.SSD,
		laptop.HDD,
		laptop.UsbC,
		laptop.UsbA,
		laptop.HDMI,
		laptop.Ethernet,
		laptop.HeadphoneJack,
		laptop.Weight,
		laptop.PriceInRupees,
		laptop.Rank,
		laptop.ReleaseDate,
	)

	if err != nil {
		Logger.Errorf("Error updating laptop, %s", err)
		return err
	}

	return nil
}

func (repo *Repo) DeleteLaptopById(ctx context.Context, id int) error {
	Logger := logger.CreateLoggerWithCtx(ctx)
	_, err := repo.db.DB().Exec(
		`DELETE FROM laptop WHERE id=$1`,
		id,
	)

	if err != nil {
		Logger.Errorf("Error deleting Laptop, %s", err)
		return err
	}

	return nil
}

func (repo *Repo) GetLaptopById(ctx context.Context, id int) (models.Laptop, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)
	var laptop models.Laptop
	sqlRow := repo.db.DB().QueryRowx(
		`SELECT * FROM laptop WHERE id=$1`,
		id,
	)

	err := sqlRow.StructScan(&laptop)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return laptop, errors.New("laptop not found")
		}
		Logger.Errorf("Error scanning fetched Laptop record, %s", err)
		return laptop, err
	}

	return laptop, err
}

func (repo *Repo) GetLaptopsByCPUId(ctx context.Context, cpu_id int) ([]models.Laptop, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)
	var laptops []models.Laptop

	sqlRows, err := repo.db.DB().Queryx(
		`SELECT * FROM laptop WHERE cpu_id=$1`,
		cpu_id)
	if err != nil {
		Logger.Errorf("Error fetching laptops, %s", err)
		return laptops, err
	}

	for sqlRows.Next() {
		var laptop models.Laptop
		err = sqlRows.StructScan(&laptop)
		if err != nil {
			Logger.Errorf("Error scanning fetched Laptop record, %s", err)
			return laptops, err
		}

		laptops = append(laptops, laptop)
	}

	return laptops, nil
}

func (repo *Repo) GetLaptopsByGPUId(ctx context.Context, gpu_id int) ([]models.Laptop, error) {
	Logger := logger.CreateLoggerWithCtx(ctx)
	var laptops []models.Laptop

	sqlRows, err := repo.db.DB().Queryx(
		`SELECT * FROM laptop WHERE gpu_id=$1`,
		gpu_id)
	if err != nil {
		Logger.Errorf("Error fetching laptops, %s", err)
		return laptops, err
	}

	for sqlRows.Next() {
		var laptop models.Laptop
		err = sqlRows.StructScan(&laptop)
		if err != nil {
			Logger.Errorf("Error scanning fetched Laptop record, %s", err)
			return laptops, err
		}

		laptops = append(laptops, laptop)
	}

	return laptops, nil
}
