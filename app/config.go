package app

type Config struct {
	GithubKey string `json:"githubkey"`
}

func NewConfig() *Config {
	return &Config{}
}

