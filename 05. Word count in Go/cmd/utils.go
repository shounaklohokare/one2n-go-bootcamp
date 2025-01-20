package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/spf13/pflag"
)

type TextFileCounts struct {
	lineCount      int
	wordCount      int
	characterCount int
}

func getTextFileCounts(fileName string) (TextFileCounts, error) {

	textFileCounts := TextFileCounts{
		lineCount:      0,
		wordCount:      0,
		characterCount: 0,
	}

	scanner, file, err := getScanner(fileName)
	if err != nil {
		return textFileCounts, err
	}

	defer file.Close()

	for scanner.Scan() {
		textFileCounts.lineCount += 1

		line := scanner.Text()
		textFileCounts.wordCount += len(strings.Fields(line))
		textFileCounts.characterCount += utf8.RuneCountInString(line)
	}

	return textFileCounts, nil
}

func getScanner(fileName string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil, file, fmt.Errorf("./wc: %v: open: No such file or directory", fileName)
	} else if os.IsPermission(err) {
		return nil, file, fmt.Errorf("./wc: %v: open: Permission denied", fileName)
	} else if err != nil {
		return nil, file, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, file, err
	}

	isDir := fileInfo.IsDir()
	if isDir {
		return nil, file, fmt.Errorf("./wc: %v: Is a directory", fileName)
	}

	return bufio.NewScanner(file), file, nil

}

func printOutput(textFileCounts TextFileCounts, fileName string, fs *pflag.FlagSet) {

	output := ""

	l := isFlagSet(fs, "linecount")
	if l {
		output += fmt.Sprintf("\t%v", textFileCounts.lineCount)
	}

	w := isFlagSet(fs, "wordcount")
	if w {
		output += fmt.Sprintf("\t%v", textFileCounts.wordCount)
	}

	c := isFlagSet(fs, "charactercount")
	if c {
		output += fmt.Sprintf("\t%v", textFileCounts.characterCount)
	}

	if !c && !w && !l {
		output += fmt.Sprintf("\t%v\t%v\t%v", textFileCounts.lineCount, textFileCounts.wordCount, textFileCounts.characterCount)
	}

	output += fmt.Sprintf("\t%v", fileName)

	fmt.Fprint(os.Stdout, output)

}

func isFlagSet(fs *pflag.FlagSet, key string) bool {

	isSet, err := fs.GetBool(key)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}

	return isSet

}
