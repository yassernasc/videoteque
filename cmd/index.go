package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"videoteque/fs"
	"videoteque/lang"
	"videoteque/movie"
	"videoteque/net"
	"videoteque/server"
	"videoteque/subtitle"
)

type serverConfig struct {
	Port int
}

type osConfig struct {
	Username string
	Password string
}

type config struct {
	Language      string
	QrCode        bool
	Server        serverConfig
	Subtitle      string
	OpenSubtitles osConfig `mapstructure:"open-subtitles"`
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
	rootCmd.Flags().StringP("language", "l", "", "language code")
	rootCmd.Flags().BoolP("qrcode", "q", false, "show qr code that redirects to the settings page")
	rootCmd.Flags().IntP("port", "p", 1200, "server port")

	viper.BindPFlags(rootCmd.Flags())                              // all flat configs
	viper.BindPFlag("Server.Port", rootCmd.Flags().Lookup("port")) // nested configs
	viper.SetDefault("language", "en")
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

func run(cmd *cobra.Command, args []string) {
	movie.Init(args[0])

	subtitle.Entry = c.Subtitle
	if c.OpenSubtitles.Username != "" {
		subtitle.SetCredentials(c.OpenSubtitles.Username, c.OpenSubtitles.Password)
	}

	server.Port = c.Server.Port
	server.ShowQrCode = c.QrCode

	lang.UserLang = c.Language
}

func validateArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return errors.New("please provide a movie entry")
	}

	_, err := movie.GetFormat(args[0])
	return err
}

func validateFlags(cmd *cobra.Command, args []string) error {
	err := validateSubtitleEntry(c.Subtitle)
	if err != nil {
		return err
	}

	err = validateOpenSubtitlesEntry(c.OpenSubtitles)
	if err != nil {
		return err
	}

	err = validatePortEntry(c.Server.Port)
	if err != nil {
		return err
	}

	return nil
}

func validateSubtitleEntry(s string) error {
	if s == "" {
		return nil
	}

	if !fs.IsFile(s) || !subtitle.IsValidEntry(s) {
		return errors.New("invalid subtitle")
	}

	return nil
}

func validatePortEntry(p int) error {
	if net.IsPortAvailable(p) {
		return nil
	}

	return errors.New("invalid port")
}

func validateOpenSubtitlesEntry(os osConfig) error {
	if os.Password == "" && os.Username == "" {
		return nil
	}

	if os.Password == "" || os.Username == "" {
		return errors.New("incomplete open subtitles info on config file")
	}

	return nil
}
