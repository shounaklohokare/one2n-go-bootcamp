package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "wc",
	Short: "A command line program that implements Unix wc like functionality.",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			return
		}

		for _, arg := range args {

			receivedCounts, err := getTextFileCounts(arg)
			if err != nil {
				fmt.Fprint(os.Stdout, err)
				os.Exit(1)
			}

			printOutput(receivedCounts, arg, cmd.Flags())
		}

	},
}

func init() {
	RootCmd.Flags().BoolP("linecount", "l", false, "flag for getting line count")
	RootCmd.Flags().BoolP("wordcount", "w", false, "flag for getting word count")
	RootCmd.Flags().BoolP("charactercount", "c", false, "flag for getting character count")
}
