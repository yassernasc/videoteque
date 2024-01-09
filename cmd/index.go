package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"lugosi/fs"
	"lugosi/movie"
	"lugosi/server"
	"lugosi/storage"
	"lugosi/subtitle"
)

var subtitlePath string
var showQrCode bool

var root = &cobra.Command{
	Use:          "lugosi",
	Short:        "Goth Movie Theater",
	Args:         validateArgs,
	PreRunE:      validateFlags,
	Run:          run,
	SilenceUsage: true,
}

func init() {
	root.Flags().StringVarP(&subtitlePath, "subtitle", "s", "", "subtitle path")
	root.Flags().BoolVarP(&showQrCode, "qrcode", "q", false, "show qrcode that links to the settings page")
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a movie entry")
	}

	_, err := movie.GetFormat(args[0])
	return err
}

func validateFlags(cmd *cobra.Command, args []string) error {
	if subtitlePath == "" {
		return nil
	}

	if !fs.IsFile(subtitlePath) || !subtitle.IsValidFile(subtitlePath) {
		return errors.New("invalid subtitle")
	}

	return nil
}

func run(cmd *cobra.Command, args []string) {
	m := movie.New(args[0])

	storage.SetMovie(m)
	storage.SetSubtitle(subtitlePath)
	storage.SetShowQrCode(showQrCode)

	server.Init()
}

func Init() {
	root.Execute()
}
