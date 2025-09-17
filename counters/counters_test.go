package counters_test

import (
	"fmt"
	"log"
	"notsecret808/ccwc/counters"
	"notsecret808/ccwc/utils"
	"testing"
)

func TestCountBytes(t *testing.T) {
	pwd, pwdError := utils.GetModuleRootDirectory()

	if pwdError != nil {
		log.Fatal("Current directory does not exist")
	}

	assetPath := fmt.Sprintf("%s/assets/%s", pwd, "test.txt")
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
