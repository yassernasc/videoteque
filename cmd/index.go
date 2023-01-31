package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"lugosi/judgment"
	"lugosi/server"
	"lugosi/storage"
)

var root = &cobra.Command{
	Use:   "lugosi",
	Short: "Goth Movie Theater",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Requires a movie")
		}
		if len(args) > 1 {
			return errors.New("Too much blood at once")
		}

		movie := args[0]
		if !judgment.IsUrl(movie) && !judgment.IsFile(movie) {
			return errors.New("Invalid movie entry")
		} else {
			return nil
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		storage.Movie = args[0]
		server.Init()
	},
	SilenceUsage: true,
}

func Init() {
	root.Execute()
}
