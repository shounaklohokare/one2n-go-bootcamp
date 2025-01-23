package main

import (
	"bufio"
	"fmt"
	"os"
)

func getScanner(fileName string) (*bufio.Scanner, *os.File, error) {

	var scanner *bufio.Scanner

	if len(fileName) == 0 {
		scanner = bufio.NewScanner(os.Stdin)
		scanner.Buffer(make([]byte, size), size)

		return scanner, nil, nil
	}

	file, err := os.Open(fileName)
	if os.IsNotExist(err) {
		return nil, file, fmt.Errorf("./mygrep: %v: open: No such file or directory", fileName)
	} else if os.IsPermission(err) {
		return nil, file, fmt.Errorf("./mygrep: %v: open: Permission denied", fileName)
	} else if err != nil {
		return nil, file, err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return nil, file, err
	}

	isDir := fileInfo.IsDir()
	if isDir {
		return nil, file, fmt.Errorf("./mygrep: %v: Is a directory", fileName)
	}

	scanner = bufio.NewScanner(file)
	scanner.Buffer(make([]byte, size), size)

	return scanner, file, nil

}
