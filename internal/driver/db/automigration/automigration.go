package automigration

import (
	"fmt"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/db"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AutoMigration interface {
	Run()
	SetDb(*gorm.DB)
}

type automigration struct {
	Db       *gorm.DB
	DbModels *[]interface{}
}

func Init(cfg *config.Configuration) {
	var mgConfigurations = map[string]AutoMigration{}
	for _, v := range cfg.Databases {
		if !v.DBAutoMigrate {
			continue
		}

		mgConfigurations[v.DBName] = &automigration{
			DbModels: &[]interface{}{
				&entity.RoleModel{},
				&entity.UserModel{},
				&entity.UserProfileModel{},
			},
		}
	}

	for k, v := range mgConfigurations {
		dbConnection, err := db.GetConnection(k)
		if err != nil {
			logrus.Error(fmt.Sprintf("failed to run automigration, database not found %s", k))
		} else {
			v.SetDb(dbConnection)
			v.Run()
			logrus.Info(fmt.Sprintf("successfully run automigration for database %s", k))
		}
	}
}

func (m *automigration) Run() {
	m.Db.AutoMigrate(*m.DbModels...)
}

func (m *automigration) SetDb(db *gorm.DB) {
	m.Db = db
}
