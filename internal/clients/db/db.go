package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	Connect(ctx context.Context) error
	Disconnect() error
	DB() *sqlx.DB
}
