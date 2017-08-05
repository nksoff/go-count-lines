package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	fileName, err := getFileName()

	if err != nil {
		logError(err)
		os.Exit(1)
	}

	lines, err := countLinesInFile(fileName)
	if err != nil {
		logError(err)
		os.Exit(2)
	}

	fmt.Printf("%d %s", lines, fileName)
}

func logError(e error) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e.Error())
}

func getFileName() (string, error) {
	if len(os.Args) == 1 {
		return "", errors.New(
			fmt.Sprintf("No file arg provided\nUsage: %s " +
				"[file]\n", os.Args[0]))
	}

	return os.Args[1], nil
}

func countLinesInFile(fileName string) (int, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return 0, err
	}

	buf := make([]byte, 1024)
	lines := 0

	for {
		readBytes, err := file.Read(buf)

		if (err != nil) {
			if readBytes == 0 && err == io.EOF {
				err = nil
			}
			return lines, err
		}

		lines += bytes.Count(buf[:readBytes], []byte{'\n'})
	}

	return lines, nil
}
