package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"

	"thl.codes/html2go/internal/commands"
)

var log = zerolog.New(zerolog.NewConsoleWriter())

var dir = "."
var logonly = false

var rootCmd = &cobra.Command{
	Use:     "html2go",
	Short:   "Convert html to go",
	Long:    `Convert html to go with gompontents`,
	Version: "0.0.1-alpha",
	Run: func(cmd *cobra.Command, args []string) {
		var mainCmd = commands.NewMainCommand(log, dir, logonly)
		if err := mainCmd.Run(context.Background(), args...); err != nil {
			cobra.CheckErr(err.Error())
		}
	},
}

func main() {
	verbose := false
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "verbose log")
	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "output dir")
	rootCmd.Flags().BoolVarP(&logonly, "logonly", "l", false, "only log generated code to stdout")

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
