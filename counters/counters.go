package counters

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func CountBytes(path string) (int, error) {
	fileInfo, error := os.Stat(path)

	if error != nil {
		return 0, fs.ErrNotExist
	}

	return int(fileInfo.Size()), nil
}

func CountChars(path string) (int, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return 0, err
	}

	fileScanner := bufio.NewScanner(readFile)

	counter := 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		counter += strings.Count(text, "")
	}

	return counter, nil
}

func CountLines(path string) (int, error) {
	readFile, err := os.Open(path)

	if err != nil {
		return 0, err
	}

	fileScanner := bufio.NewScanner(readFile)

	counter := 0

	for fileScanner.Scan() {
		counter++
	}

	return counter, nil
}

func CountWords(path string) (int, error) {
	fmt.Printf("countWords %s \n", path)
	return 0, nil
}
