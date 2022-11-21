package filemanager

import (
	"fmt"
	"io"
	"os"

	"github.com/mhkarimi1383/simple-store/internal/config"
	"github.com/mhkarimi1383/simple-store/internal/pathhelper"
	"github.com/mhkarimi1383/simple-store/types"
)

var cfg types.Config

func init() {
	cfg = config.GetConfig()
}

func SaveFile(dir, filename string, source io.Reader) error {
	fullPath := fmt.Sprintf("%v/%v", cfg.BasePath, dir)
	fullFilename := fmt.Sprintf("%v/%v", fullPath, filename)
	err := pathhelper.CreatePath(fullPath)
	if err != nil {
		return err
	}
	dst, err := os.Create(fullFilename)
	if err != nil {
		return err
	}
	if _, err = io.Copy(dst, source); err != nil {
		return err
	}
	return nil
}

func RemoveFile(dir, filename string) error {
	fullPath := fmt.Sprintf("%v/%v", cfg.BasePath, dir)
	fullFilename := fmt.Sprintf("%v/%v", fullPath, filename)
	return os.Remove(fullFilename)
}

func GetFile(dir, filename string) (io.Reader, error) {
	fullPath := fmt.Sprintf("%v/%v", cfg.BasePath, dir)
	fullFilename := fmt.Sprintf("%v/%v", fullPath, filename)
	return os.Open(fullFilename)
}
