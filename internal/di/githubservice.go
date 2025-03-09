package di

import (
	"fmt"

	github "github.com/sadhakbj/rag-with-go-ollama/internal/services"
)

func (c *Container) GithubService() *github.GithubService {
	if c.cache.GithubService != nil {
		fmt.Println("Returning cached GithubService")
		return c.cache.GithubService
	}

	githubToken := c.config.GithubToken
	httpClient := c.HTTPClient()
	fmt.Println("Creating new GithubService")
	githubService := github.NewGithubServiceWithClient(githubToken, httpClient)
	c.cache.GithubService = githubService

	return githubService
}
