package factory

import (
	"context"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/db"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/db/automigration"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/db/seeder"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/firestore"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/driver/ws"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/repository"
	"github.com/rochmanramadhani/go-lazisnu-api/internal/factory/usecase"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/gracefull"
)

type Factory struct {
	Repository repository.Factory
	Usecase    usecase.Factory
	WsHub      *ws.Hub
}

func Init(cfg *config.Configuration) (Factory, gracefull.ProcessStopper, error) {
	var stoppers []gracefull.ProcessStopper
	stopper := func(ctx context.Context) error {
		for _, st := range stoppers {
			err := st(ctx)
			if err != nil {
				return err
			}
		}
		return nil
	}

	f := Factory{}

	// db
	stopperDb, err := db.Init(cfg)
	if err != nil {
		panic(err)
	}
	stoppers = append(stoppers, stopperDb)
	conn, err := db.GetConnection(cfg.App.Connection)
	if err != nil {
		return f, stopper, err
	}

	// migration
	automigration.Init(cfg)

	// seeder
	seeder.Init(cfg)

	// ws
	f.WsHub = ws.NewHub()

	// firestore
	stopperFs, err := firestore.Init()
	stoppers = append(stoppers, stopperFs)
	fsClient := firestore.GetFirestoreClient()
	if err != nil {
		return f, stopper, err
	}

	// repository
	f.Repository = repository.Init(cfg, conn, fsClient)

	// usecase
	f.Usecase = usecase.Init(cfg, f.Repository)

	return f, stopper, nil
}
