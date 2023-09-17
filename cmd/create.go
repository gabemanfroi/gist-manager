package cmd

import (
	"github.com/gabemanfroi/gistManager/app"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Gist on GitHub",
	Long: `Create a new Gist on GitHub with the specified content, description, 
file extension, and filename. A Gist is a simple way to share code snippets or 
small files with others. For example:

To create a Gist with a code snippet:
  myapp create --content "func main() { fmt.Println('Hello, World!') }"

To create a Gist with a description:
  myapp create --content "Some important data" --description "Data file"

To specify a file extension and filename:
  myapp create --content "Code goes here" --extension "go" --filename "mycode.go"
`,

	Run: func(cmd *cobra.Command, args []string) {
		app.CreateGist(cmd)
	},
}

func init() {

	var content string
	createCmd.Flags().StringVarP(&content, "content", "c", "", "Gist content")
	createCmd.Flags().StringP("description", "d", "", "A description for the gist")
	createCmd.Flags().StringP("extension", "e", "", "A file extension for the gist")
	createCmd.Flags().StringP("filename", "f", "", "A filename for the gist")

	createCmd.Flags().SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		return pflag.NormalizedName(name)
	})

	rootCmd.AddCommand(createCmd)
}
