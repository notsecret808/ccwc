package counters

import (
	"fmt"
	"os"
)

type FileNotExistError struct {
	path string
}

func (e *FileNotExistError) Error() string {
	return fmt.Sprintf("File does not exits: %s", e.path)
}

func CountBytes(path string) (int, *FileNotExistError) {
	fileInfo, error := os.Stat(path)

	if error != nil {
		fileNotExistErr := &FileNotExistError{path}
		return 0, fileNotExistErr
	}

	return int(fileInfo.Size()), nil
}

func CountChars(path string) int {
	fmt.Printf("countChars %s \n", path)
	return 0
}

func CountLines(path string) int {
	fmt.Printf("countLines %s \n", path)
	return 0
}

func CountWords(path string) int {
	fmt.Printf("countWords %s \n", path)
	return 0
}
