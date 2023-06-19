package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"xnose/pkg/settings"
	"xnose/pkg/utils"

	"github.com/go-git/go-git/v5"
)

type RepoService interface {
	CloneRepo(ctx context.Context, url string) error
	FindFilesByExtension(ctx context.Context) ([]Result, error)
	WriteResponseToFile(ctx context.Context, results []Result) error
}

type Repo struct {
	RepoService RepoService
	Settings    *settings.Settings
}

type Result struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func NewRepoService(settings *settings.Settings) *Repo {
	return &Repo{
		Settings: settings,
	}
}

func (r *Repo) CloneRepo(ctx context.Context, name, url string) error {
	path := fmt.Sprintf("%s%s", r.Settings.StoragePath, name)
	_, err := git.PlainClone(path, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	return err
}

func (r *Repo) FindFilesByExtension(ctx context.Context) ([]Result, error) {
	var results []Result
	err := filepath.Walk(r.Settings.StoragePath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return err
		}
		if filepath.Ext(path) == r.Settings.LookUpExtension {
			results = append(results, Result{
				Name: utils.FileNameWithoutExtension(info.Name()),
				Path: path,
			})
		}
		return nil
	})
	return results, err
}

func (r *Repo) WriteResponseToFile(ctx context.Context, results []Result) error {
	data, err := json.Marshal(results)
	if err != nil {
		return err
	}
	return os.WriteFile(r.Settings.OutputPath, data, os.ModePerm)
}
