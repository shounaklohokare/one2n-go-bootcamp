package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

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

		// if len(args) == 2 {

		// 	out, err := getStringOccurencesInAFile(args[0], args[1])
		// 	if err != nil {
		// 		fmt.Fprint(os.Stdout, err)
		// 		os.Exit(1)
		// 	}

		// 	for _, line := range out {
		// 		fmt.Fprint(os.Stdout, line+"\n")
		// 	}

		// }

	},
}

func grep(r io.Reader, searchString string) ([]string, error) {

	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, size), size)

	pattern, err := regexp.Compile(searchString)
	if err != nil {
		return nil, err
	}

	var out []string
	for scanner.Scan() {
		if pattern.MatchString(scanner.Text()) {
			out = append(out, scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return out, nil
}
