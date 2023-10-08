package configs

import (
	"sync"

	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/store/sqlite"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/ilyakaznacheev/cleanenv"
)

var (
	BuildVersion string = "N/A"
	BuildDate    string = "N/A"
)

type ServerGRPC struct {
	Host       string `env:"GRPC_HOST" env-default:"localhost"`
	Port       uint   `env:"GRPC_PORT" env-default:"3200"`
	EnabledTLS bool   `env:"GRPC_ENABLED_TLS" env-default:"false"`
}

// File configuration
type File struct {
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
}

type Config struct {
	ServerGRPC ServerGRPC
	Log        logger.Options
	Auth       auth.Options
	File       File
	SQLite     sqlite.Options
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
