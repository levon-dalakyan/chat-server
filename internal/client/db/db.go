package db

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Client interface {
	DB() DB
	Close() error
}

type Query struct {
	Name     string
	QueryRaw string
}

type SQLExecer interface {
	NamedExecer
	QueryExecer
}

type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest any, q Query, args ...any) error
	ScanAllContext(ctx context.Context, dest any, q Query, args ...any) error
}

type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...any) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...any) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...any) pgx.Row
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type DB interface {
	SQLExecer
	Pinger
	Close()
}
