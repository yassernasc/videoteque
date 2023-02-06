package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"lugosi/judgment"
	"lugosi/server"
	"lugosi/storage"
)

var subtitle string

var root = &cobra.Command{
	Use:          "lugosi",
	Short:        "Goth Movie Theater",
	Args:         validateArgs,
	PreRunE:      validateFlags,
	Run:          run,
	SilenceUsage: true,
}

func init() {
	root.Flags().StringVarP(&subtitle, "subtitle", "s", "", "subtitle path")
}

func validateArgs(cmd *cobra.Command, args []string) error {
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
}

func validateFlags(cmd *cobra.Command, args []string) error {
	if subtitle == "" {
		return nil
	}

	if !judgment.IsFile(subtitle) || !judgment.IsSubtitle(subtitle) {
		return errors.New("Invalid subtitle")
	}

	return nil
}

func run(cmd *cobra.Command, args []string) {
	storage.Movie = args[0]
	if subtitle != "" {
		storage.Subtitle = subtitle
	}
	server.Init()
}

func Init() {
	root.Execute()
}
