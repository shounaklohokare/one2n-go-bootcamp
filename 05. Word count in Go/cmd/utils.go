package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func getLineCount(fileName string) (int, error) {

	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return -1, fmt.Errorf("./wc: %v: open: No such file or directory", fileName)
	} else if os.IsPermission(err) {
		return -1, fmt.Errorf("./wc: %v: open: Permission denied", fileName)
	} else if err != nil {
		return -1, err
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return -1, err
	}

	isDir := fileInfo.IsDir()
	if isDir {
		return -1, fmt.Errorf("./wc: %v: Is a directory", fileName)
	}

	scanner := bufio.NewScanner(file)

	lineCount := 0
	for scanner.Scan() {
		lineCount += 1
	}

	return lineCount, nil
}
