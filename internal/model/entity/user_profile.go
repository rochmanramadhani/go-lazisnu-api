package entity

type UserProfileEntity struct {
	Address  *string `gorm:"not null"`
	Phone    *string `gorm:"not null"`
	FilePath string  `gorm:"not null"`

	// fk
	UserID uint64 `gorm:"not null"`
}

type UserProfileModel struct {
	Entity
	UserProfileEntity
}

func (UserProfileModel) TableName() string {
	return "user_profiles"
}
