package counters_test

import (
	"fmt"
	"notsecret808/ccwc/counters"
	"notsecret808/ccwc/utils"
	"testing"
)

func geTestFilePath() string {
	pwd, pwdError := utils.GetModuleRootDirectory()

	if pwdError != nil {
		panic(pwdError)
	}

	assetPath := fmt.Sprintf("%s/assets/%s", pwd, "test.txt")

	return assetPath
}

func TestCountBytes(t *testing.T) {
	assetPath := geTestFilePath()

	bytesCount, error := counters.CountBytes(assetPath)

	if error != nil {
		t.Error(error)
		return
	}

	if bytesCount != 342190 {
		message := fmt.Sprintf("Bytes count does not match %d", bytesCount)
		t.Error(message)
	}
}

func TestCountLines(t *testing.T) {
	assetPath := geTestFilePath()

	linesCount, error := counters.CountLines(assetPath)

	if error != nil {
		t.Error(error)
		return
	}

	if linesCount != 7145 {
		message := fmt.Sprintf("Lines count does not match %d", linesCount)
		t.Error(message)
	}
}

func TestCountWords(t *testing.T) {
	assetPath := geTestFilePath()

	wordsCount, error := counters.CountWords(assetPath)

	if error != nil {
		t.Error(error)
		return
	}

	if wordsCount != 58164 {
		message := fmt.Sprintf("Lines count does not match %d", wordsCount)
		t.Error(message)
	}
}
