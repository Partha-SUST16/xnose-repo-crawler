package main

import (
	"context"
	"xnose/app"
	"xnose/pkg/settings"
)

func main() {
	settings := settings.NewSettings()
	repo := app.NewRepoService(settings)
	repo.CloneRepo(context.TODO(), "https://github.com/go-git/go-git")
}
