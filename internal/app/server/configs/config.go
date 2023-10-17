package configs

import (
	"sync"

	"github.com/Orendev/gokeeper/internal/app/server/delivery/grpc"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/ilyakaznacheev/cleanenv"
)

type Delivery struct {
	GRPC grpc.Options
}

// DB basic settings postgres
type DB struct {
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

type Config struct {
	Postgres DB
	Delivery Delivery
	Log      logger.Options
	Auth     auth.Options
}

var configInstance *Config
var configErr error

// New constructor a new instance of Config.
func New() (*Config, error) {
	if configInstance == nil {
		var readConfigOnce sync.Once

		readConfigOnce.Do(func() {
			configInstance = &Config{}
			configErr = cleanenv.ReadConfig(".env", configInstance)
		})
	}

	return configInstance, configErr
}
