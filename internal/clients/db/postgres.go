package db

import (
	"context"
	"cquest/internal/logger"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"time"
)

type PostgresDB struct {
	Sqlx     *sqlx.DB
	Host     string
	Port     string
	Username string
	Password string
	SSLMode  string
	Database string
}

func NewPostgresDB(host string, port string, username string, password string, sslMode string, database string) *PostgresDB {
	return &PostgresDB{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		SSLMode:  sslMode,
		Database: database,
	}
}

func (p *PostgresDB) DB() *sqlx.DB { return p.Sqlx }
func (p *PostgresDB) Connect(ctx context.Context) error {
	var err error
	var count int8
	Logger := logger.CreateLoggerWithCtx(ctx)

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", p.Username, p.Password, p.Host, p.Port, p.Database, p.SSLMode)

	for {
		p.Sqlx, err = sqlx.Connect("postgres", dbUrl)
		if err != nil {
			Logger.Errorf("Error connecting to Postgres: %v", err)
			count++
		} else {
			Logger.Infof("connected to postgres %s", dbUrl)
			p.Sqlx.SetMaxOpenConns(5000)
			p.Sqlx.SetMaxIdleConns(1000)
			p.Sqlx.SetConnMaxLifetime(2 * time.Minute)
			break
		}

		if count > 5 {
			Logger.Errorf(err.Error())
			return err
		}
		Logger.Warnf("Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)

	}

	return nil
}

func (p *PostgresDB) Disconnect() error { return p.Sqlx.Close() }
