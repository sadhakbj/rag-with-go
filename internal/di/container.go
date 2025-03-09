package di

import (
	"sync"

	"github.com/sadhakbj/rag-with-go-ollama/internal/config"
	github "github.com/sadhakbj/rag-with-go-ollama/internal/services"
	"github.com/sadhakbj/rag-with-go-ollama/internal/utils/httpclient"
)

type Container struct {
	config     *config.Config
	cache      *cache
	mu         sync.Mutex
	httpClient httpclient.HTTPClient
}

type cache struct {
	GithubService *github.GithubService
}

func NewContainer(config *config.Config) *Container {
	return &Container{
		config:     config,
		cache:      &cache{},
		httpClient: httpclient.NewClient(),
	}
}
