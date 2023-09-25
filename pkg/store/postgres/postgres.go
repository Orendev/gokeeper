package postgres

import (
	"context"
	"fmt"
	"github.com/Orendev/gokeeper/internal/app/server/configs"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
)

// Store - structure describing the Postgres.
type Store struct {
	Pool *pgxpool.Pool
}

// ENV to pg params
// PGHOST
// PGPORT
// PGDATABASE
// PGUSER
// PGPASSWORD
// PGPAPPNAME
// PGSSLMODE

// New - constructor a new instance of Postgres.
func New(ctx context.Context, cfg configs.DB) (*Store, error) {

	config, err := pgxpool.ParseConfig(toDSN(cfg))
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	err = conn.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		Pool: conn,
	}, nil
}

func toDSN(cfg configs.DB) string {
	var args []string
	if len(cfg.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", cfg.Host))
	}
	if cfg.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", cfg.Port))
	}

	if len(cfg.Name) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", cfg.Name))
	}

	if len(cfg.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", cfg.User))
	}

	if len(cfg.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", cfg.Password))
	}

	if len(cfg.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", cfg.SSLMode))
	}

	return strings.Join(args, " ")
}
