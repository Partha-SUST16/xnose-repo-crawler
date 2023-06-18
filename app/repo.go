package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"xnose/pkg/settings"

	"github.com/go-git/go-git/v5"
)

type RepoService interface {
	CloneRepo(ctx context.Context, url string)
}

type Repo struct {
	RepoService RepoService
	Settings    *settings.Settings
}

func NewRepoService(settings *settings.Settings) *Repo {
	return &Repo{
		Settings: settings,
	}
}

func (r *Repo) CloneRepo(ctx context.Context, name, url string) {
	path := fmt.Sprintf("%s%s", r.Settings.StoragePath, name)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	log.Println(err)
}
