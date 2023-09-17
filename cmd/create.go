package cmd

import (
	"github.com/gabemanfroi/gistManager/app"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new resource",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

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
