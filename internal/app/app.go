package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/sadhakbj/rag-with-go-ollama/internal/config"
	"github.com/sadhakbj/rag-with-go-ollama/internal/di"
	"github.com/sadhakbj/rag-with-go-ollama/internal/utils/logger"
)

type App struct {
	Name      string
	Version   string
	Config    *config.Config
	Container *di.Container
	Logger    *slog.Logger
}

func NewApp(config *config.Config) *App {
	return &App{
		Name:    config.AppName,
		Version: config.AppVersion,
		Config:  config,
		Logger:  logger.InitializeLogger(config.AppName, false, slog.LevelInfo),
	}
}

func (a *App) Run() {
	a.Container = di.NewContainer(a.Config)
	a.Logger.Info("Starting application", "version", a.Version)
	githubService := a.Container.GithubService()

	context := context.Background()
	prs, err := githubService.ListPRs(context, "sadhakbj", "rag-with-laravel-ollama")
	if err != nil {
		a.Logger.Error("Failed to list PRs", "error", err)
		os.Exit(1)
	}

	for _, v := range prs {
		fmt.Printf("PR: %d, Title: %s, State: %s\n", v.Number, v.Title, v.State)
	}

	githubService2 := a.Container.GithubService()

	pr2s, err := githubService2.ListPRs(context, "sadhakbj", "rag-with-laravel-ollama")
	if err != nil {
		a.Logger.Error("Failed to list PRs second time", "error", err)
		log.Fatalf("failed to list PRs second time: %v", err)
		os.Exit(1)
	}

	a.Logger.Info("Listing PRs")

	for _, v := range pr2s {
		fmt.Printf("PR: %d, Title: %s, State: %s\n", v.Number, v.Title, v.State)
	}
}
