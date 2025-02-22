package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env string `env:"ENV" envDefault:"dev"`
	App AppConfig
	DB  DBConfig
}

type AppConfig struct {
	Host string `env:"APP_HOST" envDefault:"localhost"`
	Port string `env:"APP_PORT" envDefault:"8080"`
}

type DBConfig struct {
	User     string `env:"POSTGRES_USER" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	DB       string `env:"POSTGRES_DB" env-required:"true"`
}

func DSN(cfg *DBConfig) string {
	return fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s",
		cfg.Port, cfg.User, cfg.Password, cfg.DB)
}

func MustLoad() *Config {
	var cfg Config
	var err error

	configPath := fetchConfigPath()
	if configPath != "" {
		err = godotenv.Load(configPath)
	} else {
		err = godotenv.Load()
	}

	if err != nil {
		log.Println("Не удалось загрузить конфигурацию")
	}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		panic("Конфигурация некорректна:" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "Путь к файлу конфигурации")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
