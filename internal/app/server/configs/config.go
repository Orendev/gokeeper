package configs

import (
	"sync"

	"github.com/Orendev/gokeeper/internal/app/server/delivery/grpc"
	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/Orendev/gokeeper/pkg/store/postgres"
	"github.com/Orendev/gokeeper/pkg/tools/auth"
	"github.com/ilyakaznacheev/cleanenv"
)

type Delivery struct {
	GRPC grpc.Options
}

type Config struct {
	Postgres postgres.Options
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
