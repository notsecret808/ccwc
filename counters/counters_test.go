package counters_test

import (
	"fmt"
	"io"
	"notsecret808/ccwc/counters"
	"notsecret808/ccwc/utils"
	"os"
	"path/filepath"
	"testing"
)

func geTestFilePath() string {
	pwd, pwdError := utils.GetModuleRootDirectory()

	if pwdError != nil {
		panic(pwdError)
	}

	assetPath := filepath.Join(pwd, "assets", "test.txt")
	return assetPath
}

func getAsset(t *testing.T) io.Reader {
	assetPath := geTestFilePath()

	readAsset, err := os.Open(assetPath)

	if err != nil {
		t.Error(err)
	}

	return readAsset
}

func TestCountBytes(t *testing.T) {

	bytesCount, error := counters.CountBytes(getAsset(t))

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
	linesCount, error := counters.CountLines(getAsset(t))

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
	wordsCount, error := counters.CountWords(getAsset(t))

	if error != nil {
		t.Error(error)
		return
	}

	if wordsCount != 58164 {
		message := fmt.Sprintf("Lines count does not match %d", wordsCount)
		t.Error(message)
	}
}
