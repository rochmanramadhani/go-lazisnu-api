package config

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/configor"
	"github.com/sirupsen/logrus"
)

type Configuration struct {
	App        App        `json:"app"`
	Swagger    Swagger    `json:"swagger"`
	JWT        JWT        `json:"jwt"`
	Databases  []Database `json:"databases"`
	Connection Connection `json:"connection"`
	Driver     Driver     `json:"driver"`
	Asset      Asset      `json:"asset"`
	Queue      Queue      `json:"queue"`
}

var Config *Configuration = &Configuration{}

func LoadWithYml(path string) (*Configuration, error) {
	if path == "" {
		wd, err := os.Getwd()
		if err != nil {
			return Config, err
		}
		path = fmt.Sprintf("%s/config/config.%s.yml", wd, os.Getenv("ENV"))
	}

	if !isValidEnvironment() {
		path = fmt.Sprintf("/run/secrets/%s", os.Getenv("CONFIG"))
	}

	err := configor.New(&configor.Config{AutoReload: true, AutoReloadInterval: time.Minute}).Load(Config, path)
	if err != nil {
		logrus.Info(err)
		return Config, err
	}

	return Config, nil
}

func LoadWithEnv() (*Configuration, error) {
	dbLogLevel, err := strconv.Atoi(os.Getenv("DB_LOG_LEVEL"))
	if err != nil {
		return nil, err
	}
	queueImageMaxLine, err := strconv.Atoi(os.Getenv("QUEUE_IMAGE_MAX_LINE"))
	if err != nil {
		return nil, err
	}
	queueFileMaxLine, err := strconv.Atoi(os.Getenv("QUEUE_FILE_MAX_LINE"))
	if err != nil {
		return nil, err
	}

	Config.App.Name = os.Getenv("APP_NAME")
	Config.App.Key = os.Getenv("APP_KEY")
	Config.App.Port = os.Getenv("APP_PORT")
	Config.App.Host = os.Getenv("APP_HOST")
	Config.App.Version = os.Getenv("APP_VERSION")
	Config.App.Connection = os.Getenv("APP_CONNECTION")
	Config = &Configuration{
		App: App{
			Name:       os.Getenv("APP_NAME"),
			Key:        os.Getenv("APP_KEY"),
			Port:       os.Getenv("APP_PORT"),
			Host:       os.Getenv("APP_HOST"),
			Version:    os.Getenv("APP_VERSION"),
			Connection: os.Getenv("APP_CONNECTION"),
		},
		Swagger: Swagger{
			SwaggerHost:   os.Getenv("SWAGGER_HOST"),
			SwaggerScheme: os.Getenv("SWAGGER_SCHEME"),
			SwaggerPrefix: os.Getenv("SWAGGER_PREFIX"),
		},
		JWT: JWT{
			Secret:        os.Getenv("JWT_SECRET"),
			RefreshSecret: os.Getenv("JWT_REFRESH_SECRET"),
		},
		Databases: []Database{
			{
				DBHost:        os.Getenv("DB_HOST"),
				DBUser:        os.Getenv("DB_USER"),
				DBPass:        os.Getenv("DB_PASS"),
				DBPort:        os.Getenv("DB_PORT"),
				DBName:        os.Getenv("DB_NAME"),
				DBProvider:    os.Getenv("DB_PROVIDER"),
				DBSSL:         os.Getenv("DB_SSL"),
				DBTZ:          os.Getenv("DB_TZ"),
				DBAutoMigrate: os.Getenv("DB_AUTO_MIGRATE") == "true",
				DBSeeder:      os.Getenv("DB_SEEDER") == "true",
				DBLogLevel:    dbLogLevel,
			},
		},
		Connection: Connection{
			Primary: os.Getenv("CONNECTION_PRIMARY"),
			Replica: os.Getenv("CONNECTION_REPLICA"),
		},
		Driver: Driver{
			Cron: Cron{
				Enabled: os.Getenv("DRIVER_CRON_ENABLED") == "true",
			},
			Firestore: Firestore{
				Credentials: os.Getenv("DRIVER_FIRESTORE_CREDENTIALS"),
				ProjectID:   os.Getenv("DRIVER_FIRESTORE_PROJECT_ID"),
			},
			Elasticsearch: Elasticsearch{
				Url:      os.Getenv("DRIVER_ELASTICSEARCH_URL"),
				User:     os.Getenv("DRIVER_ELASTICSEARCH_USER"),
				Password: os.Getenv("DRIVER_ELASTICSEARCH_PASSWORD"),
			},
			Sentry: Sentry{
				Dsn: os.Getenv("DRIVER_SENTRY_DSN"),
			},
		},
		Asset: Asset{
			ImageTempDir:   os.Getenv("IMAGE_TEMP_DIR"),
			ImageExtension: os.Getenv("IMAGE_EXTENSION"),
			ImageQuality:   os.Getenv("IMAGE_QUALITY"),
			FileTempDir:    os.Getenv("FILE_TEMP_DIR"),
			FileExtension:  os.Getenv("FILE_EXTENSION"),
		},
		Queue: Queue{
			QueueImageMaxLine: queueImageMaxLine,
			QueueFileMaxLine:  queueFileMaxLine,
		},
	}

	return Config, nil
}

func (Configuration) String() string {
	sb := strings.Builder{}
	return sb.String()
}

func (c Configuration) Raw() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func isValidEnvironment() bool {
	allowedEnvironments := []string{"development", "debug", "local"}
	currentEnv := os.Getenv("ENV")

	for _, env := range allowedEnvironments {
		if currentEnv == env {
			return true
		}
	}

	return false
}
