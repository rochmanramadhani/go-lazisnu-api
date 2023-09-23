package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rochmanramadhani/go-lazisnu-api/docs"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	api "github.com/rochmanramadhani/go-lazisnu-api/internal/delivery/api"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/delivery/api/middleware"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/gracefull"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init(cfg *config.Configuration, f factory.Factory) (gracefull.ProcessStarter, gracefull.ProcessStopper) {
	var (
		APP        = cfg.App.Name
		VERSION    = cfg.App.Version
		DOC_HOST   = cfg.Swagger.SwaggerHost
		DOC_SCHEME = cfg.Swagger.SwaggerScheme
	)
	// echo
	e := echo.New()

	// Serve static files from the "/public" folder
	e.Static("/public", "public")

	// index
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to %s version %s", APP, VERSION)
		return c.String(http.StatusOK, message)
	})

	// doc
	docs.SwaggerInfo.Title = APP
	docs.SwaggerInfo.Version = VERSION
	docs.SwaggerInfo.Host = DOC_HOST
	docs.SwaggerInfo.Schemes = []string{DOC_SCHEME}
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// delivery
	middleware.Init(e)
	api.Init(e, f)
	// ws.Init(e, f)

	return func() error {
			return e.Start(":" + cfg.App.Port)
		}, func(ctx context.Context) error {
			return e.Shutdown(ctx)
		}
}
