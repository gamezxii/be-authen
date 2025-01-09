package config

import (
	"log"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Server *Server
		Db     *Db
		Amqp   *Amqp
		Aws    *Aws
	}

	Server struct {
		Port int
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
		MinPool  int
		MaxPool  int
	}

	Amqp struct {
		Url string
	}

	Aws struct {
		AccessKey string
		SecretKey string
		Region    string
	}
)

var (
	once           sync.Once
	configInstance *Config
)

func GetConfig() *Config {
	once.Do(func() {
		viper.SetConfigFile(".env")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		if err := viper.ReadInConfig(); err != nil {
			log.Printf("No .env file found or error reading .env file: %v", err)
		}

		configInstance = &Config{
			Server: &Server{
				Port: viper.GetInt("SERVER_PORT"),
			},
			Db: &Db{
				Host:     viper.GetString("DB_HOST"),
				Port:     viper.GetInt("DB_PORT"),
				User:     viper.GetString("DB_USER"),
				Password: viper.GetString("DB_PASSWORD"),
				DBName:   viper.GetString("DB_NAME"),
				SSLMode:  viper.GetString("DB_SSLMODE"),
				TimeZone: viper.GetString("DB_TIMEZONE"),
				MinPool:  viper.GetInt("DB_TIMEZONE"),
				MaxPool:  viper.GetInt("DB_TIMEZONE"),
			},
			Amqp: &Amqp{
				Url: viper.GetString("AMQP_URL"),
			},
			Aws: &Aws{
				AccessKey: viper.GetString("AWS_ACCESS_KEY_ID"),
				SecretKey: viper.GetString("AWS_SECRET_ACCESS_KEY"),
				Region:    viper.GetString("AWS_REGION"),
			},
		}

		log.Printf("Loaded config: %+v", configInstance)

		// ตรวจสอบค่าที่โหลดมาไม่เป็น nil
		if configInstance.Server == nil {
			log.Fatal("Server config is missing")
		}
		if configInstance.Db == nil {
			log.Fatal("Database config is missing")
		}
		if configInstance.Amqp == nil {
			log.Fatal("AMQP config is missing")
		}

	})

	return configInstance
}
