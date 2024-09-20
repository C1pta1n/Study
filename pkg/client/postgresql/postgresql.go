package postgresql

import (
	"context"

	"github.com/jackc/pgx"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgx.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

// func NewClient(ctx context.Context, sc config.Config) (pool *pgx.Conn, err error) {
// 	// pool, err = pgx.Connect(ctx, 5*time.Second))
// }
