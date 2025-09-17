package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func ParseCmdParams() ([]string, string, error) {
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
