package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func getLineCount(fileName string) (int, error) {

	scanner, file, err := getScanner(fileName)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	lineCount := 0
	for scanner.Scan() {
		lineCount += 1
	}

	return lineCount, nil
}

func getWordCount(fileName string) (int, error) {

	scanner, file, err := getScanner(fileName)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	wordCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		wordCount += len(strings.Fields(line))
	}

	return wordCount, nil

}

func getCharacterCount(fileName string) (int, error) {

	scanner, file, err := getScanner(fileName)
	if err != nil {
		return -1, err
	}

	defer file.Close()

	characterCount := 0
	for scanner.Scan() {
		line := scanner.Text()

		characterCount += utf8.RuneCountInString(line)
	}

	return characterCount, nil

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
