package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	//Debug bool `yaml:"debug" env:"DEBUG_ENABLED" env-default:"true"`

	HTTPPort     string        `yaml:"http_port" env:"HTTP_PORT" env-default:"8082"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"30s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"30s"`
	MaxHeaderMB  int           `yaml:"max_header_mb" env:"MAX_HEADER_MB" env-default:"20"`

	//GRPCPort string `yaml:"grpc_port" env:"GRPC_PORT" env-default:"9082"`

	RedisTimeout  time.Duration `yaml:"redis_timeout" env:"REDIS_TIMEOUT" env-default:"30s"`
	RedisAddr     string        `yaml:"redis_addr" env:"REDIS_ADDR" env-default:"localhost:6379"`
	RedisPassword string        `yaml:"redis_password" env:"REDIS_PASSWORD" env-default:""`
}

func Init() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
