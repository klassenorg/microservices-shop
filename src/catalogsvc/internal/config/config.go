package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"time"
)

type Config struct {
	HTTPPort     string        `yaml:"http_port" env:"HTTP_PORT" env-default:"8081"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"30s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"30s"`
	MaxHeaderMB  int           `yaml:"max_header_mb" env:"MAX_HEADER_MB" env-default:"20"`

	GRPCPort string `yaml:"grpc_port" env:"GRPC_PORT" env-default:"9081"`

	MongoTimeout  time.Duration `yaml:"mongo_timeout" env:"MONGO_TIMEOUT" env-default:"30s"`
	MongoURI      string        `yaml:"mongo_uri" env:"MONGO_URI" env-default:"mongodb://localhost:27017"`
	MongoUser     string        `yaml:"mongo_user" env:"MONGO_USER" env-default:""`
	MongoPassword string        `yaml:"mongo_password" env:"MONGO_PASSWORD" env-default:""`
	MongoDBName   string        `yaml:"mongo_db_name" env:"MONGO_DB_NAME" env-default:"catalog"`
}

func Init() *Config {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
