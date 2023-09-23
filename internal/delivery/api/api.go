package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/delivery/api/handler"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
)

func Init(e *echo.Echo, f factory.Factory) {
	// routes
	prefix := "api/v1"

	handler.NewRole(f).Route(e.Group(prefix + "/roles"))
	handler.NewUser(f).Route(e.Group(prefix + "/users"))
	handler.NewAuth(f).Route(e.Group(prefix + "/auth"))

	handler.NewFaq(f).Route(e.Group(prefix + "/faqs"))
	handler.NewFaqCategory(f).Route(e.Group(prefix + "/faq-categories"))

	handler.NewDonationType(f).Route(e.Group(prefix + "/donation-types"))
}
