package github

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sadhakbj/rag-with-go-ollama/internal/utils/httpclient"
)

type GithubService struct {
	httpClient httpclient.HTTPClient
	baseURL    string
	token      string
}

func NewGithubService(token string) *GithubService {
	return &GithubService{
		httpClient: httpclient.NewClient(),
		baseURL:    "https://api.github.com",
		token:      token,
	}
}

// PR represents a GitHub pull request
type PR struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
	State  string `json:"state"`
}

func (g *GithubService) ListPRs(ctx context.Context, owner string, repo string) ([]PR, error) {
	url := fmt.Sprintf("%s/repos/%s/%s/pulls", g.baseURL, owner, repo)

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", g.token),
		"Accept":        "application/vnd.github.v3+json",
	}

	res, err := g.httpClient.Get(ctx, url, headers)
	if err != nil {
		return nil, fmt.Errorf("failed to get PRs: %w", err)
	}

	var prs []PR
	if err := json.Unmarshal(res, &prs); err != nil {
		return nil, fmt.Errorf("failed to unmarshal PRs: %w", err)
	}

	return prs, nil
}
