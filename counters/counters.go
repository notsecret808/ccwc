package counters

import (
	"fmt"
	"notsecret808/ccwc/errors"
	"os"
)

func CountBytes(path string) (int, *errors.FileNotExistError) {
	fileInfo, error := os.Stat(path)

	if error != nil {
		fileNotExistErr := &errors.FileNotExistError{Path: path}
		return 0, fileNotExistErr
	}

	return int(fileInfo.Size()), nil
}

func CountChars(path string) (int, *errors.FileNotExistError) {
	fmt.Printf("countChars %s \n", path)
	return 0, nil
}

func CountLines(path string) (int, *errors.FileNotExistError) {
	fmt.Printf("countLines %s \n", path)
	return 0, nil
}

func CountWords(path string) (int, *errors.FileNotExistError) {
	fmt.Printf("countWords %s \n", path)
	return 0, nil
}
