package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TextFileCounts struct {
	lineCount      int
	wordCount      int
	characterCount int
}

const size = 1024 * 1024

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

	if len(fileName) != 0 {
		defer file.Close()
	}

	for scanner.Scan() {
		line := scanner.Text()

		textFileCounts.lineCount += 1
		textFileCounts.wordCount += len(strings.Fields(line))
		textFileCounts.characterCount += len(line) + 1
	}

	return textFileCounts, nil
}

func getScanner(fileName string) (*bufio.Scanner, *os.File, error) {

	var scanner *bufio.Scanner

	if len(fileName) == 0 {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Buffer(make([]byte, size), size)

		return scanner, nil, nil
	}

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

	scanner = bufio.NewScanner(file)
	scanner.Buffer(make([]byte, size), size)

	return scanner, file, nil

}

func printOutput(textFileCounts TextFileCounts, fileName string, wcFlags WcFlags) {

	output := ""

	l := wcFlags.lineFlag
	if l {
		output += fmt.Sprintf("\t%v", textFileCounts.lineCount)
	}

	w := wcFlags.wordFlag
	if w {
		output += fmt.Sprintf("\t%v", textFileCounts.wordCount)
	}

	c := wcFlags.characterFlag
	if c {
		output += fmt.Sprintf("\t%v", textFileCounts.characterCount)
	}

	if !c && !w && !l {
		output += fmt.Sprintf("\t%v\t%v\t%v", textFileCounts.lineCount, textFileCounts.wordCount, textFileCounts.characterCount)
	}

	output += fmt.Sprintf("\t%v", fileName)

	fmt.Fprint(os.Stdout, output)

}
