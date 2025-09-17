package counters

import (
	"fmt"
	"io/fs"
	"os"
)

func CountBytes(path string) (int, error) {
	fileInfo, error := os.Stat(path)

	if error != nil {
		return 0, fs.ErrNotExist
	}

	return int(fileInfo.Size()), nil
}

func CountChars(path string) (int, error) {

	return 0, nil
}

func CountLines(path string) (int, error) {
	fmt.Printf("countLines %s \n", path)
	return 0, nil
}

func CountWords(path string) (int, error) {
	fmt.Printf("countWords %s \n", path)
	return 0, nil
}
