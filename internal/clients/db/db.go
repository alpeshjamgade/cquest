package db

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type DB interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	DB() *sqlx.DB
}
