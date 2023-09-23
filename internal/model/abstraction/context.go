package abstraction

import (
	"gorm.io/gorm"
)

type AuthContext struct {
	UserID    uint64
	RoleID    uint64
	CompanyID uint64
	Name      string
}

type TrxContext struct {
	Db *gorm.DB
}

type UploadFileContext struct {
	FilePath string
	FileName string
}
