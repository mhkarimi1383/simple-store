package pathhelper

import (
	"errors"
	"os"
)

// CreatePath create a directory if not exist
func CreatePath(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return err
}
