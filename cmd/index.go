package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
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
		fmt.Println(args[0])
	},
}

func Init() {
	root.Execute()
}
