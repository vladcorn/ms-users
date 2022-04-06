package conf

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	APP  APP
	DB   DB
	GRPC GRPC
	HTTP HTTP
	LOG  LOG
}

// APP describes
type APP struct {
}

// DB describes database config
type DB struct {
	Name     string `env:"DB_NAME,required"`
	Host     string `env:"DB_HOST,required"`
	Port     int    `env:"DB_PORT,required"`
	User     string `env:"DB_USER,required"`
	Password string `env:"DB_PASSWORD,required"`
	SSL      string `env:"DB_SSL,required"`
	TimeZone string `env:"DB_TIMEZONE,required"`

	OpenLimit int           `env:"DB_OPEN_LIMIT" envDefault:"5"`
	IdleLimit int           `env:"DB_IDLE_LIMIT" envDefault:"5"`
	ConnLife  time.Duration `env:"DB_CONN_LIFE" envDefault:"5m"`
}

type GRPC struct {
	Host string `env:"GRPC_HOST" envDefault:"0.0.0.0"`
	Port string `env:"GRPC_PORT" envDefault:"10000"`
}

func (grpc *GRPC) Addr() string {
	return grpc.Host + ":" + grpc.Port
}

func (d *DB) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s timezone=%s", d.Host, d.Port, d.User, d.Password, d.Name, d.SSL, d.TimeZone)
}

type HTTP struct {
	Host string   `env:"HTTP_HOST" envDefault:"0.0.0.0"`
	Port string   `env:"HTTP_PORT" envDefault:"11000"`
	CORS []string `env:"HTTP_CORS" envSeparator:";" envDefault:"*"`
}

func (http *HTTP) Addr() string {
	return http.Host + ":" + http.Port
}

type LOG struct {
	Level      string `env:"LOG_LEVEL" envDefault:"debug"`
	Json       bool   `env:"LOG_JSON" envDefault:"false"`
	TimeFormat string `env:"LOG_TIME_FORMAT" envDefault:"2006-01-02T15:04:05Z"`
}

var once sync.Once

func New() Config {
	var c Config
	once.Do(func() {
		if err := env.Parse(&c); err != nil {
			panic("parsing configuration")
		}
		c.print()
	})
	return c
}

func (cfg *Config) print() {
	jsonConfig, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonConfig))
}
