package di

import (
	"github.com/sadhakbj/rag-with-go-ollama/internal/config"
	github "github.com/sadhakbj/rag-with-go-ollama/internal/services"
	"github.com/sadhakbj/rag-with-go-ollama/internal/utils/httpclient"
)

type Container struct {
	config *config.Config
	cache  *cache
}

type cache struct {
	GithubService *github.GithubService
	HTTPClient    httpclient.HTTPClient
}

func NewContainer(config *config.Config) *Container {
	return &Container{
		config: config,
		cache:  &cache{},
	}
}
