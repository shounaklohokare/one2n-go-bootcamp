package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var size int = 1024 * 1024

func main() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stdout, err)
	}

}

var RootCmd = cobra.Command{
	Use:   "grep",
	Short: "A commnad line program that implements Unix grep like functionality",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 2 {

			out, err := getStringOccurencesInAFile(args[0], args[1])
			if err != nil {
				fmt.Fprint(os.Stdout, err)
				os.Exit(1)
			}

			for _, line := range out {
				fmt.Fprint(os.Stdout, line+"\n")
			}

		}

	},
}

func getStringOccurencesInAFile(searchString, fileName string) ([]string, error) {

	scanner, file, err := getScanner(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var out []string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchString) {
			out = append(out, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return out, nil

}

func getStringOccurencesFromStdin(searchString string) ([]string, error) {

	scanner, _, err := getScanner("")
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}

	var out []string
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, searchString) {
			out = append(out, line)
		}
	}

	return out, nil

}
