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

		fileName, _ := cmd.Flags().GetString("linecount")

		lineCount, err := getLineCount(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("\t%d %v\n", lineCount, fileName)

	},
}

func init() {
	RootCmd.Flags().StringP("linecount", "l", "file342.txt", "file name for getting line count")
}
