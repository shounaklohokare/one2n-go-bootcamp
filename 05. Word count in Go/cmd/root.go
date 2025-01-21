package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type WcFlags struct {
	lineFlag      bool
	wordFlag      bool
	characterFlag bool
}

var wcFlags WcFlags

var RootCmd = &cobra.Command{
	Use:   "wc",
	Short: "A command line program that implements Unix wc like functionality.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			args = []string{""} // passing empty string in case of reading input from stdin
		}

		for _, arg := range args {

			receivedCounts, err := getTextFileCounts(arg)
			if err != nil {
				fmt.Fprint(os.Stdout, err)
				os.Exit(1)
			}

			printOutput(receivedCounts, arg, wcFlags)
		}

	},
}

func init() {
	RootCmd.Flags().BoolVarP(&wcFlags.lineFlag, "linecount", "l", false, "flag for getting line count")
	RootCmd.Flags().BoolVarP(&wcFlags.wordFlag, "wordcount", "w", false, "flag for getting word count")
	RootCmd.Flags().BoolVarP(&wcFlags.characterFlag, "charactercount", "c", false, "flag for getting character count")
}
