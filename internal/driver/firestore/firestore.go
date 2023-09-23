package firestore

import (
	"context"
	"sync"

	"github.com/rochmanramadhani/go-lazisnu-api/internal/config"
	"github.com/rochmanramadhani/go-lazisnu-api/pkg/util/gracefull"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var (
	firestoreClient *firestore.Client
	firestoreOnce   sync.Once
	stopper         = gracefull.NilStopper
)

func Init() (gracefull.ProcessStopper, error) {
	return stopper, nil
	ctx := context.Background()
	opt := option.WithCredentialsFile(config.Config.Driver.Firestore.Credentials)
	config := &firebase.Config{ProjectID: config.Config.Driver.Firestore.ProjectID}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return stopper, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return stopper, err
	}

	firestoreOnce.Do(func() {
		firestoreClient = client
		stopper = func(ctx context.Context) error {
			return firestoreClient.Close()
		}
	})

	return stopper, nil
}

func GetFirestoreClient() *firestore.Client {
	return firestoreClient
}
