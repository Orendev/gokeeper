package sqlite

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

type Store struct {
	mu sync.Mutex
	DB *sql.DB
}

const File string = "./services/keeperClient/store.db"

// Options basic settings postgres
type Options struct {
	Timeout        uint64 `env:"SQLITE_TIMEOUT"`
	DaraSourceName string `env:"SQLITE_DATA_SOURCE_NAME" env-default:"./services/keeperClient/store.db"`
	DefaultLimit   uint64 `env:"DEFAULT_LIMIT" env-default:"10"`
	DefaultOffset  uint64 `env:"DEFAULT_OFFSET" env-default:"10"`
	MigrationsDir  string `env:"SQLITE_MIGRATIONS_DIR" env-default:"./services/keeperClient/internal/repository/storage/sqlite/migrations"`
}

func New(ctx context.Context, o Options) (*Store, error) {
	db, err := sql.Open("sqlite3", o.DaraSourceName)

	if err != nil {
		return nil, err
	}

	return &Store{DB: db}, nil
}
