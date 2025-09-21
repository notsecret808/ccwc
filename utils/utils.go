package utils

import (
	"os/exec"
	"strings"
)

func GetModuleRootDirectory() (string, error) {
	out, err := exec.Command("go", "list", "-m", "-f", "{{.Dir}}").Output()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}
