package temp

import (
	"context"
	"log"
	"multimodal-search/internal/config"

	"github.com/weaviate/weaviate-go-client/v5/weaviate"
)

type Repository interface {
	GetTest(ctx context.Context) error
}

type repository struct {
	client *weaviate.Client
}

func NewRepository(configWeaviate config.Weaviate) Repository {
	client, err := weaviate.NewClient(weaviate.Config{
		Host:   configWeaviate.Host,
		Scheme: configWeaviate.Schema,
	})
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()

	ping, err := client.Misc().LiveChecker().Do(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Weaviate connection: %v", ping)

	return &repository{client: client}
}

func (r *repository) GetTest(ctx context.Context) error {
	return nil
}
