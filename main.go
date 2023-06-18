package main

import (
	"context"
	"log"
	"xnose/app"
	"xnose/pkg/response"
	"xnose/pkg/settings"
)

func main() {
	log.SetFlags(log.Ltime)
	settings := settings.NewSettings()
	responses := response.ReadResponse(settings)
	repo := app.NewRepoService(settings)
	for _, item := range responses.Items {
		log.Println(item.CloneURL)
		repo.CloneRepo(context.Background(), item.Name, item.CloneURL)
	}
}
