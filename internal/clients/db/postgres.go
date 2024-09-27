package db

import (
	"context"
	"cquest/internal/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type PostgresDB struct {
	Sqlx     *sqlx.DB
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func NewPostgresDB(host string, port string, username string, password string, database string) *PostgresDB {
	return &PostgresDB{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
	}
}

func (p *PostgresDB) DB() *sqlx.DB { return p.Sqlx }
func (p *PostgresDB) Connect(ctx context.Context) error {
	var err error
	Logger := logger.CreateLoggerWithCtx(ctx)

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", p.Username, p.Password, p.Host, p.Port, p.Database)

	p.Sqlx, err = sqlx.Connect("postgres", dbUrl)
	if err != nil {
		Logger.Errorw("error connecting to db", "error", err)
		return err
	}
	Logger.Infof("connected to postgres %s", dbUrl)

	p.Sqlx.SetMaxOpenConns(5000)
	p.Sqlx.SetMaxIdleConns(1000)
	p.Sqlx.SetConnMaxLifetime(2 * time.Minute)

	return nil
}

func (p *PostgresDB) Disconnect() error { return p.Sqlx.Close() }
