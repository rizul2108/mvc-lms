package utils

import (
	"path/filepath"
)

func GetCurrentDirPath() (string, error) {
	return filepath.Abs(".")
}
