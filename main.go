package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	options, path, err := parseParams()

	if err != nil {
		log.Fatal(err)
	}

	stats := make([]int, 4)

	for _, option := range options {
		switch option {
		case "-c", "--bytes":
			bytesCount := countBytes(path)
			stats[0] = bytesCount
		case "-m", "--chars":
			charsCount := countChars(path)
			stats[1] = charsCount
		case "-l", "--lines":
			linesCount := countLines(path)
			stats[2] = linesCount
		case "-w", "--words":
			wordsCount := countWords(path)
			stats[3] = wordsCount
		default:
			log.Fatal("Incorrect option")
		}
	}

	fmt.Printf("%d %d %d %s \n", stats[0], stats[1], stats[2], path)
}

func parseParams() ([]string, string, error) {
	args := os.Args[1:]

	if len(args) < 1 {
		return os.Args, "", errors.New("options are not provided")
	}

	options := args[:len(args)-1]
	filePath := args[len(args)-1]

	_, absErr := os.Stat(filePath)

	currentDir, dirErr := os.Getwd()

	if dirErr != nil {
		log.Fatal("Cannot get current dir")
	}

	relativePath := fmt.Sprintf("%s/%s", currentDir, filePath)
	_, relErr := os.Stat(relativePath)

	if errors.Is(absErr, os.ErrNotExist) || errors.Is(relErr, os.ErrNotExist) {
		return args, filePath, os.ErrNotExist
	}
	return options, filePath, nil
}

func countBytes(path string) int {
	fmt.Printf("countBytes: %s \n", path)
	return 0
}

func countChars(path string) int {
	fmt.Printf("countChars %s \n", path)
	return 0
}

func countLines(path string) int {
	fmt.Printf("countLines %s \n", path)
	return 0
}

func countWords(path string) int {
	fmt.Printf("countWords %s \n", path)
	return 0
}
