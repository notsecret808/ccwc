package cmdParser

import (
	"log"
	"os"
	"regexp"
)

func isCommandOption(option string) bool {
	result, _ := regexp.MatchString(`^--?\S+$`, option)
	return result

}

func validateFile(filePath string) (path string, fileError error) {
	file, err := os.Open(filePath)

	if err != nil {
		return filePath, err
	}

	defer file.Close()

	return filePath, nil
}

func distinctCommandParamsByCategories() (commandTargets []string, commandOptions []string) {
	args := os.Args[1:]
	filePaths := make([]string, 0)
	options := make([]string, 0)

	for _, option := range args {
		if !isCommandOption(option) {
			filePaths = append(filePaths, option)
		} else {
			options = append(options, option)
		}
	}

	return filePaths, options
}

func filterFilesByExistence(commandTargets []string) (fileNames []string) {
	fileNames = make([]string, 0)

	for _, fileName := range commandTargets {
		path, err := validateFile(fileName)

		if err != nil {
			log.Println(err)
			return
		}

		fileNames = append(fileNames, path)
	}

	return
}

func ParseCommandParams() (filePaths []string, options []string) {
	filePaths, options = distinctCommandParamsByCategories()

	if len(filePaths) > 0 {
		filePaths = filterFilesByExistence(filePaths)
	}

	return
}
