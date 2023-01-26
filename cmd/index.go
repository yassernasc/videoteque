package cmd

import (
	"errors"
	"github.com/spf13/cobra"
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
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		storage.Movie = args[0]
		server.Init()
	},
}

func Init() {
	root.Execute()
}
