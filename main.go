package main

import (
	"github.com/joho/godotenv"
	"sync"
	"time"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/cron"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/http"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/sentry"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/gracefull"
)

// @title LAZISNU API
// @version 0.0.1
// @description This is a LAZISNU API server.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8801/rest-api/v-1
// @contact.name Rochman Ramadhani Chiefto Irawan
// @contact.url rochmanramadhani12@gmail.com
// @contact.email https://www.linkedin.com/in/rochmanramadhan/
// @BasePath /

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// if you want to load from env file
	cfg, err := config.LoadWithEnv()
	if err != nil {
		panic(err)
	}

	// if you want to load from yml
	//cfg, err := config.LoadWithYml("")
	//if err != nil {
	//	panic(err)
	//}

	// factory
	wg := new(sync.WaitGroup)
	f, stopperFactory, err := factory.Init(cfg)
	if err != nil {
		panic(err)
	}

	sentry.Init()
	starterApi, stopperApi := http.Init(cfg, f)
	cron.Init(cfg, f)

	wg.Add(1)
	go func() {
		gracefull.StartProcessAtBackground(starterApi)
		gracefull.StopProcessAtBackground(time.Second*10, stopperApi, stopperFactory)
		wg.Done()
	}()

	wg.Wait()
}
