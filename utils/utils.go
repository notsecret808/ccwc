package utils

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ParseCmdParams() ([]string, string, error) {
	fi, _ := os.Stdin.Stat()
	args := os.Args[1:]

	if len(args) < 1 {
		return os.Args, "", errors.New("options are not provided")
	}

	if (fi.Mode() & os.ModeCharDevice) == 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

	}

	options := args[:len(args)-1]
	filePath := args[len(args)-1]

	_, absErr := os.Stat(filePath)

	pwd, pwdErr := os.Getwd()

	if pwdErr != nil {
		log.Fatal("Cannot get current dir")
	}

	relativeFilePath := fmt.Sprintf("%s/%s", pwd, filePath)
	_, relErr := os.Stat(relativeFilePath)

	if errors.Is(relErr, os.ErrNotExist) && errors.Is(absErr, os.ErrNotExist) {
		return args, filePath, os.ErrNotExist
	}

	if relErr == nil {
		return options, relativeFilePath, nil
	} else {
		return options, filePath, nil
	}
}

func GetModuleRootDirectory() (string, error) {
	out, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}").Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
