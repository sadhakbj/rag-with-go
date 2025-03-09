package di

import (
	"fmt"

	github "github.com/sadhakbj/rag-with-go-ollama/internal/services"
)

func (c *Container) GithubService() *github.GithubService {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.cache.GithubService != nil {
		fmt.Println("Returning cached GithubService")
		return c.cache.GithubService
	}

	githubToken := c.config.GithubToken
	fmt.Println("Creating new GithubService")
	githubService := github.NewGithubServiceWithClient(githubToken, c.httpClient)
	c.cache.GithubService = githubService

	return githubService
}
