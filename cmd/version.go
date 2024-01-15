package cmd

import (
	"fmt"
	"github.com/carlmjohnson/versioninfo"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show installed version",
	Run: func(cmd *cobra.Command, args []string) {
		v := versioninfo.LastCommit.Format("2006.01.02")
		fmt.Println("vidéotequè (vt)", v)
	},
}
