package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"videoteque/fs"
	"videoteque/movie"
	"videoteque/server"
	"videoteque/storage"
	"videoteque/subtitle"
)

type config struct {
	Port     int
	QrCode   bool
	Subtitle string
}

var (
	c       config
	cfgFile string
	rootCmd = &cobra.Command{
		Use:          "vt <movie-entry>",
		Short:        "tool for watching movies",
		Args:         validateArgs,
		PreRunE:      validateFlags,
		Run:          run,
		SilenceUsage: true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "config file")
	rootCmd.Flags().StringP("subtitle", "s", "", "subtitle path")
	rootCmd.Flags().BoolP("qrcode", "q", false, "show qrcode to access the settings page")
	rootCmd.Flags().IntP("port", "p", 1200, "server port")

	viper.BindPFlags(rootCmd.Flags())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigType("toml")
		viper.SetConfigName("config")
		viper.AddConfigPath(fs.ConfigDir())
	}

	viper.ReadInConfig()
	viper.Unmarshal(&c)
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a movie entry")
	}

	_, err := movie.GetFormat(args[0])
	return err
}

func validateFlags(cmd *cobra.Command, args []string) error {
	s := c.Subtitle

	if s == "" {
		return nil
	}

	if !fs.IsFile(s) || !subtitle.IsValidFile(s) {
		return errors.New("invalid subtitle")
	}

	return nil
}

func run(cmd *cobra.Command, args []string) {
	m := movie.New(args[0])

	storage.Movie = m
	storage.Subtitle = c.Subtitle
	storage.ShowQrCode = c.QrCode
	storage.Port = c.Port

	server.Init()
}
