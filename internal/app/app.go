package app

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	github "github.com/sadhakbj/rag-with-go-ollama/internal/services"
)

type App struct {
	Name    string
	Version string
}

func NewApp() *App {
	return &App{
		Name:    "Rag with Go",
		Version: "0.0.1",
	}
}

func (a *App) Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	githubToken := os.Getenv("GITHUB_TOKEN")

	githubService := github.NewGithubService(githubToken)

	context := context.Background()
	prs, err := githubService.ListPRs(context, "sadhakbj", "rag-with-laravel-ollama")
	if err != nil {
		log.Fatalf("failed to list PRs: %v", err)
	}

	for _, v := range prs {
		fmt.Printf("PR: %d, Title: %s, State: %s\n", v.Number, v.Title, v.State)
	}
}
