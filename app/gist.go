package app

import (
	"context"
	"github.com/gabemanfroi/gistManager/internal"
	"github.com/google/go-github/v55/github"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

type Gist struct {
	Description string
	Content     string
	Extension   *string
	Filename    *string
}

func CreateGist(cmd *cobra.Command) {

	gist := promptGist(cmd)
	PersistToGit(gist)
}

func promptGist(cmd *cobra.Command) Gist {
	description := getDescription(cmd)
	content := getContent(cmd)
	extension := getExtension(cmd)
	filename := getFileName(cmd, description)

	gist := Gist{
		Description: strings.Title(description),
		Content:     content,
		Extension:   &extension,
		Filename:    &filename,
	}
	return gist
}

func getFileName(cmd *cobra.Command, description string) string {
	filename, _ := cmd.Flags().GetString("filename")
	if filename == "" {
		filename = internal.GetInput(internal.Content{
			Label: "Enter the filename of the gist",
		}, 0)
	}
	if filename == "" {
		filename = strings.Title(description)
	}
	return filename
}

func getExtension(cmd *cobra.Command) string {
	extension, _ := cmd.Flags().GetString("extension")
	if extension == "" {
		extension = internal.GetInput(internal.Content{
			Label: "Enter the extension of the gist",
		}, 0)
	}
	return extension
}

func getDescription(cmd *cobra.Command) string {
	description, _ := cmd.Flags().GetString("description")
	if description == "" {
		description = internal.GetInput(internal.Content{
			Label:    "Enter the description of the gist",
			ErrorMsg: "Description cannot be empty",
		}, 5)
	}
	return description
}

func getContent(cmd *cobra.Command) string {
	content, _ := cmd.Flags().GetString("content")
	if content == "" {
		content = internal.GetMultiline(internal.Content{
			Label:    "Enter the content of the gist",
			ErrorMsg: "Content cannot be empty",
		})
	}
	return content
}

func PersistToGit(gist Gist) {
	_, _, err := GetClient().Gists.Create(context.Background(), &github.Gist{
		Description: github.String(strings.Title(gist.Description)),
		Public:      github.Bool(true),
		Files: map[github.GistFilename]github.GistFile{
			github.GistFilename(*gist.Filename + *gist.Extension): {
				Content: github.String(gist.Content),
			},
		},
	})
	if err != nil {
		log.Fatal(err.Error())
	}
}
