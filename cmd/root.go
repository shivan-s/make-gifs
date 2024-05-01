package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "makegif",
	Short: "makegif turns movies files into gifs.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Echo: ", strings.Join(args, " "))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
