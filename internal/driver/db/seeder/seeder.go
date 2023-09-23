package seeder

import (
	"context"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/db"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/db/seeder/seed"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/ctxval"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/trxmanager"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Seeder interface {
	Run() error
	SetDb(*gorm.DB)
}

type seeder struct {
	Db    *gorm.DB
	Seeds []seed.Seed
}

type seederEntity struct {
	Tag string `gorm:"index:seeder_tag,unique"`
}

func (seederEntity) TableName() string {
	return `seeder_migrations`
}

func Init(cfg *config.Configuration) {
	var mgConfigurations = map[string]Seeder{}
	for _, v := range cfg.Databases {
		if !v.DBSeeder {
			continue
		}

		mgConfigurations[v.DBName] = &seeder{
			Seeds: []seed.Seed{
				&seed.RoleSeed{},
				&seed.UserSeed{},
			},
		}
	}

	for k, v := range mgConfigurations {
		dbConnection, err := db.GetConnection(k)
		if err != nil {
			logrus.Error("failed to run seeder, database not found", k)
		} else {
			v.SetDb(dbConnection)
			if err := v.Run(); err != nil {
				break
			}

			logrus.Info("successfully run seeder for database", k)
		}
	}
}

func (m *seeder) Run() error {
	if !m.Db.Migrator().HasTable(seederEntity.TableName(seederEntity{})) {
		err := m.Db.Migrator().CreateTable(&seederEntity{})
		if err != nil {
			logrus.Error("failed to run seeder, create seed entity ", err.Error())
			return err
		}
	}

	for _, v := range m.Seeds {
		if err := trxmanager.New(m.Db).WithTrx(context.Background(), func(ctx context.Context) error {
			trx := ctxval.GetTrxValue(ctx)
			seed := seederEntity{
				Tag: v.GetTag(),
			}

			var seedExist seederEntity
			if err := trx.Where("tag", v.GetTag()).Find(&seedExist).Error; err != nil {
				return err
			}
			if seedExist.Tag != "" {
				logrus.Info("skip seed, cause has been executed ", v.GetTag())
				return nil
			}

			if err := trx.Create(&seed).Error; err != nil {
				return err
			}

			if err := v.Run(m.Db); err != nil {
				return err
			}

			return nil
		}); err != nil {
			logrus.Error("skip seed, cause some error in ", v.GetTag(), err.Error())
			continue
		}
	}

	return nil
}

func (m *seeder) SetDb(db *gorm.DB) {
	m.Db = db
}
