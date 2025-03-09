package app

import (
	"context"
	"fmt"
	"log"

	"github.com/sadhakbj/rag-with-go-ollama/internal/config"
	"github.com/sadhakbj/rag-with-go-ollama/internal/di"
)

type App struct {
	Name      string
	Version   string
	Container *di.Container
}

func NewApp() *App {
	return &App{
		Name:    "Rag with Go",
		Version: "0.0.1",
	}
}

func (a *App) Run() {
	cfg := config.LoadConfig()
	a.Container = di.NewContainer(cfg)

	githubService := a.Container.GithubService()

	context := context.Background()
	prs, err := githubService.ListPRs(context, "sadhakbj", "rag-with-laravel-ollama")
	if err != nil {
		log.Fatalf("failed to list PRs: %v", err)
	}

	for _, v := range prs {
		fmt.Printf("PR: %d, Title: %s, State: %s\n", v.Number, v.Title, v.State)
	}

	githubService2 := a.Container.GithubService()

	pr2s, err := githubService2.ListPRs(context, "sadhakbj", "rag-with-laravel-ollama")
	if err != nil {
		log.Fatalf("failed to list PRs second time: %v", err)
	}

	fmt.Println("Second time")

	for _, v := range pr2s {
		fmt.Printf("PR: %d, Title: %s, State: %s\n", v.Number, v.Title, v.State)
	}
}
