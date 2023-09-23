package db

import (
	"fmt"

	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/gracefull"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	dbPostgreSQL struct {
		db
		SslMode  string
		Tz       string
		LogLevel int
	}
)

func (c *dbPostgreSQL) Init() (*gorm.DB, gracefull.ProcessStopper, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", c.Host, c.User, c.Pass, c.Name, c.Port, c.SslMode, c.Tz)

	logLevel := logger.Info
	if c.LogLevel != 0 {
		logLevel = logger.LogLevel(c.LogLevel)
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, gracefull.NilStopper, err
	}
	return db, gracefull.NilStopper, nil
}
