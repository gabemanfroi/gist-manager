package app

import (
	"fmt"
	"github.com/google/go-github/v55/github"
	"os"
)

func GetClient() *github.Client {
	fmt.Println(os.Getenv("GITHUB_ACCESS_TOKEN"))
	return github.NewClient(nil).WithAuthToken(os.Getenv("GITHUB_ACCESS_TOKEN"))
}
