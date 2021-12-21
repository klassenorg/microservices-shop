package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	Debug          bool   `yaml:"debug" env:"DEBUG_ENABLED" env-default:"true"`
	DebugPprofPort string `yaml:"debug_pprof_port" env:"DEBUG_PPROF_PORT" env-default:"8885"`

	HTTPPort     string        `yaml:"http_port" env:"HTTP_PORT" env-default:"8085"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"30s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"30s"`
	MaxHeaderMB  int           `yaml:"max_header_mb" env:"MAX_HEADER_MB" env-default:"20"`

	GRPCPort string `yaml:"grpc_port" env:"GRPC_PORT" env-default:"9085"`

	CartSvcGRPCAddr    string `yaml:"cart_svc_grpc_addr" env:"CART_SVC_GRPC_PORT" env-default:"localhost:9082"`
	PricingSvcGRPCAddr string `yaml:"pricing_svc_grpc_addr" env:"PRICING_SVC_GRPC_PORT" env-default:"localhost:9083"`

	PostgresHost     string `yaml:"postgres_host" env:"POSTGRES_HOST" env-default:"localhost"`
	PostgresPort     string `yaml:"postgres_port" env:"POSTGRES_PORT" env-default:"5432"`
	PostgresUser     string `yaml:"postgres_user" env:"POSTGRES_USER" env-default:"postgres"`
	PostgresDBName   string `yaml:"postgres_db_name" env:"POSTGRES_DB_NAME" env-default:"postgres"`
	PostgresPassword string `yaml:"postgres_password" env:"POSTGRES_PASSWORD" env-default:"password"`
	PostgresSSLMode  string `yaml:"postgres_ssl_mode" env:"POSTGRES_SSL_MODE" env-default:"disable"`
}

func Init() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
