package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var size int = 1024 * 1024

func main() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Fprint(os.Stdout, err)
	}

}

type GrepResult struct {
	FileName string
	Result   []string
	Error    error
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

func writeToFile(r io.Reader, searchString, outputFile string) error {

	output, err := grep(r, searchString)
	if err != nil {
		return err
	}

	err = checkFileExists(outputFile)
	if err != nil {
		return err
	}

	content := strings.Join(output, "\n")
	err = os.WriteFile(outputFile, []byte(content), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	return nil

}

func checkFileExists(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		return fmt.Errorf("file '%s' already exists", filename)
	}
	if os.IsNotExist(err) {
		return nil
	}
	return err
}

func grepInDir(dir, searchString string) ([]GrepResult, error) {

	var grepResult []GrepResult

	fileNames, err := listFiles(dir)
	if err != nil {
		return grepResult, err
	}

	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			grepResult = append(grepResult, GrepResult{
				FileName: fileName,
				Error:    err,
			})
			continue
		}

		defer file.Close()

		result, err := grep(file, searchString)
		if err != nil {
			grepResult = append(grepResult, GrepResult{
				FileName: fileName,
				Error:    err,
			})
			continue
		}

		grepResult = append(grepResult, GrepResult{
			FileName: fileName,
			Result:   result,
		})

	}

	return grepResult, err

}

func listFiles(dir string) ([]string, error) {
	var fileNames []string

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			fileNames = append(fileNames, path)
		}

		return nil
	})

	return fileNames, err
}
