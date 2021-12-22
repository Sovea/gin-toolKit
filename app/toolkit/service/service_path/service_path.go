package service_path

import (
	"os"
	"strings"
)

func GetCurrentPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return strings.Replace(dir, "\\", "/", -1), nil
}
