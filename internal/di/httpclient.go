package di

import "github.com/sadhakbj/rag-with-go-ollama/internal/utils/httpclient"

func (c *Container) HTTPClient() httpclient.HTTPClient {
	if c.cache.HTTPClient != nil {
		return c.cache.HTTPClient
	}

	httpClient := httpclient.NewClient()
	c.cache.HTTPClient = httpClient

	return httpClient
}
