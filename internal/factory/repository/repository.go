package repository

import (
	"cloud.google.com/go/firestore"
	"firebase.google.com/go/messaging"
	el "github.com/olivere/elastic/v7"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	dbRepository "github.com/rochmanramadhani/go-lazisnu-api/internal/repository/db"
	"gorm.io/gorm"
)

type Factory struct {
	Db  *gorm.DB
	Es  *el.Client
	Fcm *messaging.Client
	Fs  *firestore.Client

	Role        dbRepository.Role
	User        dbRepository.User
	UserProfile dbRepository.UserProfile

	Faq         dbRepository.Faq
	FaqCategory dbRepository.FaqCategory

	DonationType dbRepository.DonationType
}

func Init(cfg *config.Configuration, db *gorm.DB, fs *firestore.Client) Factory {
	f := Factory{}

	// db
	f.Db = db
	f.Role = dbRepository.NewRole(f.Db)
	f.User = dbRepository.NewUser(f.Db)
	f.UserProfile = dbRepository.NewUserProfile(f.Db)

	f.Faq = dbRepository.NewFaq(f.Db)
	f.FaqCategory = dbRepository.NewFaqCategory(f.Db)

	f.DonationType = dbRepository.NewDonationType(f.Db)

	// firestore
	f.Fs = fs

	return f
}
