package postgres

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
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

// Options basic settings postgres
type Options struct {
	Host          string `env:"DB_HOST" env-default:"localhost"`
	Name          string `env:"DB_NAME" env-default:"gokeeper"`
	Port          uint16 `env:"DB_PORT" env-default:"5432"`
	User          string `env:"DB_USERNAME" env-default:"gokeeper"`
	Password      string `env:"DB_PASSWORD" env-default:"secret"`
	SSLMode       string `env:"DB_SSL_MODE" env-default:"disable"`
	DefaultLimit  uint64 `env:"DEFAULT_LIMIT" env-default:"10"`
	DefaultOffset uint64 `env:"DEFAULT_OFFSET" env-default:"10"`
	MigrationsDir string `env:"POSTGRES_MIGRATIONS_DIR" env-default:"./services/keeperServer/internal/repository/storage/postgres/migrations"`
}

// New - constructor a new instance of Postgres.
func New(ctx context.Context, cfg Options) (*Store, error) {

	config, err := pgxpool.ParseConfig(cfg.toDSN())
	if err != nil {
		return nil, err
	}

	conn, err := pgxpool.ConnectConfig(ctx, config)
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

func (cfg *Options) toDSN() string {
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
