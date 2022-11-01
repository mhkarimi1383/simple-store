package pathhelper

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func CreatePath(path string) {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
