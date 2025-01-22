package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var size int32 = 1024 * 1024

var RootCmd = cobra.Command{
	Use:   "grep",
	Short: "A commnad line program that implements Unix grep like functionality",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func getStringOccurencesInAFile(searchString, fileName string) ([]string, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	out := make([]string, 10)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchString) {
			out = append(out, line)
		}
	}

	return out, nil

}
