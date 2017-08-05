package main

import (
	"os"
	"strings"
	"testing"
)

func TestGetFileName(t *testing.T) {
	f := func(args []string, expectedFileName string, expectError bool) {
		os.Args = args
		returnedFileName, err := getFileName()

		if expectError && err == nil {
			t.Errorf("Expected error, got result string '%s'. Input args: [%s]", returnedFileName, strings.Join(args, ","))
		} else if !expectError && err != nil {
			t.Errorf("Not expected error, got '%s'. Input args: [%s]", err, strings.Join(args, ","))
		} else if returnedFileName != expectedFileName {
			t.Errorf("Expected '%s', found '%s'", expectedFileName, returnedFileName)
		}
	}

	f([]string{"./count_lines", "123.txt"}, "123.txt", false)
	f([]string{"./count_lines", "some/long/long/long/long/long/long/long/long/long/long/file.txt"}, "some/long/long/long/long/long/long/long/long/long/long/file.txt", false)

	f([]string{"./count_lines", "123.txt", "one_more_file.txt"}, "123.txt", false)

	f([]string{"./count_lines"}, "", true)
}

func TestCountLinesInFile(t *testing.T) {
	f := func(fileName string, expectedLines int, expectError bool) {
		lines, err := countLinesInFile(fileName)

		if expectError && err == nil {
			t.Errorf("Expected error, got result number '%d'. Input file: %s", lines, fileName)
		} else if !expectError && err != nil {
			t.Errorf("Not expected error, got '%s'. Input file: %s", err, fileName)
		} else {
			if expectedLines != lines {
				t.Errorf("Expected %d lines, got %d in file '%s'", lines, expectedLines, fileName)
			}
		}
	}

	f("fixtures/file.txt", 3, false)
	f("fixtures/file1.txt", 7, false)
	f("fixtures/file2.txt", 48998, false)
	f("fixtures/some-unknown-file.txt", 0, true)
}
