package app

import (
	"github.com/google/go-github/v55/github"
	"os"
)

func GetClient() *github.Client {
	return github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_ACCESS_TOKEN"))
}
