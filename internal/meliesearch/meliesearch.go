package meliesearch

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/meilisearch/meilisearch-go"
)

type Meliesearch struct {
	meilisearch.ServiceManager
}

func New (url string , apiKey string) (*Meliesearch , error){
	client := meilisearch.New(url , meilisearch.WithAPIKey(apiKey))

	health , err := client.Health()
	if err != nil {
		return nil, fmt.Errorf("meliesearch: ping failed: %w", err)
	}

	slog.Info("meilisearch connection established", "status", health.Status)

	return &Meliesearch{client} , nil
}

func (ms *Meliesearch) Close(){
	// it`s a http request under the hood
}

func (ms *Meliesearch) Ping(ctx context.Context) error {
	_ , err := ms.Health()
	if err != nil {
		return fmt.Errorf("meliesearch: ping failed: %w", err)
	}
	return nil
}