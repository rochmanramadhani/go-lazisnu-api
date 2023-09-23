package usecase

import (
	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/repository"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/usecase"
)

type Factory struct {
	Role usecase.Role
	User usecase.User
	Auth usecase.Auth

	Faq         usecase.Faq
	FaqCategory usecase.FaqCategory

	DonationType usecase.DonationType
}

func Init(cfg *config.Configuration, r repository.Factory) Factory {
	f := Factory{}

	f.Role = usecase.NewRole(cfg, r)
	f.User = usecase.NewUser(cfg, r)
	f.Auth = usecase.NewAuth(cfg, r)

	f.Faq = usecase.NewFaq(cfg, r)
	f.FaqCategory = usecase.NewFaqCategory(cfg, r)

	f.DonationType = usecase.NewDonationType(cfg, r)

	return f
}
