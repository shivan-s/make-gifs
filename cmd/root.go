package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

var rootCmd = &cobra.Command{
	Use:   "makegif",
	Short: "makegif turns movies files into gifs.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			var err error
			parent := filepath.Dir(arg)
			palleteDir := filepath.Join(parent, ".temp.png")
			err = ffmpeg.Input(arg).
				Output(palleteDir, ffmpeg.KwArgs{"vf": "palettegen"}).
				OverWriteOutput().ErrorToStdOut().Run()
			if err != nil {
				panic(err)
			}
			err = ffmpeg.Input(arg).
				Output(strings.TrimSuffix(arg, filepath.Ext(arg))+".gif", ffmpeg.KwArgs{
					"i":              ".temp.png",
					"filter_complex": "fps=10,scale=1280:-1[x];[x][1:v]paletteuse"}).
				OverWriteOutput().Run()
			if err != nil {
				panic(err)
			}
			err = os.Remove(palleteDir)
			if err != nil {
				panic(err)
			}
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
