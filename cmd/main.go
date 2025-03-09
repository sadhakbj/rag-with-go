package main

import (
	"fmt"

	"github.com/sadhakbj/rag-with-go-ollama/internal/app"
	"github.com/sadhakbj/rag-with-go-ollama/internal/config"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Println(cfg.AppName)
	fmt.Println(cfg.AppVersion)
	fmt.Println(cfg.GithubToken)

	app := app.NewApp(cfg)
	app.Run()
}
