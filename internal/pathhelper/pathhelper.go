package pathhelper

import (
	"errors"
	"os"
)

// CreatePath create a directory if not exist
func CreatePath(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return os.MkdirAll(path, os.ModePerm)
	} else {
		return err
	}
}
