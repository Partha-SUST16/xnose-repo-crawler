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
	log.Println("Total Repo Count", len(responses.Items))
	repo := app.NewRepoService(settings)
	for _, item := range responses.Items {
		err := repo.CloneRepo(context.TODO(), item.Name, item.CloneURL)
		if err != nil {
			log.Println(err.Error())
		}
	}
	results, err := repo.FindFilesByExtension(context.TODO())
	if err != nil {
		log.Println(err.Error())
	}
	err = repo.WriteResponseToFile(context.TODO(), results)
	if err != nil {
		log.Println(err.Error())
	}
}
