package configs

import (
	"sync"

	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/store/sqlite"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/ilyakaznacheev/cleanenv"
)

type GRPC struct {
	Host string `env:"GRPC_HOST" env-default:"localhost"`
	Port uint   `env:"GRPC_PORT" env-default:"3200"`
}

type Config struct {
	GRPC   GRPC
	Log    logger.Options
	Auth   auth.Options
	SQLite sqlite.Options
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
