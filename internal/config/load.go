package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type (
	Config struct {
		HTTP HTTP `yaml:"http" env-prefix:"http"`
		Auth Auth `yaml:"auth" env-prefix:"auth"`
		PG   PG   `yaml:"postgres"`
	}

	Cache struct {
		Capacity uint          `yaml:"capacity"`
		TTL      time.Duration `yaml:"ttl"`
	}

	HTTP struct {
		Cache          Cache         `yaml:"cache"`
		SwaggerPort    uint          `env-required:"true" yaml:"swagger_port"`
		Port           string        `env-required:"true" yaml:"port"`
		ReadTimeout    time.Duration `env-required:"true" yaml:"read_timeout"`
		WriteTimeout   time.Duration `env-required:"true" yaml:"write_timeout"`
		IdleTimeout    time.Duration `env-required:"true" yaml:"idle_timeout"`
		MaxHeaderBytes int           `env-required:"true" yaml:"max_header_bytes"`
	}

	PG struct {
		URL             string        `env:"PG_URL" env-default:"postgres://user:password@postgres:5432/postgres?sslmode=disable"`
		MaxOpenConns    int           `env-required:"true" yaml:"max_open_conns"`
		MaxIdleConns    int           `env-required:"true" yaml:"max_idle_conns"`
		ConnMaxIdleTime time.Duration `env-required:"true" yaml:"conn_max_idle_time"`
		ConnMaxLifetime time.Duration `env-required:"true" yaml:"conn_max_lifetime"`
	}

	Auth struct {
		PasswordCostBcrypt int           `env-required:"true" yaml:"password_cost_bcrypt"`
		AccessTokenTTL     time.Duration `env-required:"true" yaml:"access_token_ttl"`
		SigningKey         string        `env:"JWT_SIGNING_KEY" env-default:"wdkadwadwakpklrbjb"`
	}
)

const (
	configPath = "./config/config.yml"
)

func NewConfig() (Config, error) {
	cfg := Config{}

	path, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		path = configPath
	}

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func MustNewConfig() Config {
	cfg, err := NewConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
