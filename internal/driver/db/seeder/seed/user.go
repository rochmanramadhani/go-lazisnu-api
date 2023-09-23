package seed

import (
	"fmt"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/model/entity"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserSeed struct{}

func (s *UserSeed) Run(conn *gorm.DB) error {
	trx := conn.Begin()

	var roles []entity.RoleModel
	if err := trx.Model(&entity.RoleModel{}).Find(&roles).Error; err != nil {
		return err
	}

	for _, role := range roles {
		email := fmt.Sprintf(`%s@gmail.com`, role.Name)
		user := entity.UserModel{
			UserEntity: entity.UserEntity{
				Name:     &role.Name,
				Email:    &email,
				Password: role.Name,
				RoleID:   &role.ID,
			},
		}
		if err := trx.Create(&user).Error; err != nil {
			trx.Rollback()
			logrus.Error(err)
			return err
		}

		address := "xxx"
		phone := "021"
		userProfile := entity.UserProfileModel{
			UserProfileEntity: entity.UserProfileEntity{
				Address: &address,
				Phone:   &phone,
				UserID:  user.ID,
			},
		}
		if err := trx.Create(&userProfile).Error; err != nil {
			trx.Rollback()
			logrus.Error(err)
			return err
		}
	}

	if err := trx.Commit().Error; err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func (s *UserSeed) GetTag() string {
	return `user_seed`
}
