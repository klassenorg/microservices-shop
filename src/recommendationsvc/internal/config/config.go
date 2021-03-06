package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Debug          bool   `yaml:"debug" env:"DEBUG_ENABLED" env-default:"true"`
	DebugPprofPort string `yaml:"debug_pprof_port" env:"DEBUG_PPROF_PORT" env-default:"8884"`

	HTTPPort     string        `yaml:"http_port" env:"HTTP_PORT" env-default:"8084"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"30s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"30s"`
	MaxHeaderMB  int           `yaml:"max_header_mb" env:"MAX_HEADER_MB" env-default:"20"`

	GRPCPort string `yaml:"grpc_port" env:"GRPC_PORT" env-default:"9084"`

	CatalogSvcGRPCAddr string `yaml:"catalog_svc_grpc_addr" env:"CATALOG_SVC_GRPC_PORT" env-default:"localhost:9081"`
}

func Init() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
