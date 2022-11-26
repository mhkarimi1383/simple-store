package filemanager

import (
	"fmt"
	"io"
	"os"

	"github.com/mhkarimi1383/simple-store/internal/config"
	"github.com/mhkarimi1383/simple-store/internal/pathhelper"
	"github.com/mhkarimi1383/simple-store/internal/typeconverters"
	"github.com/mhkarimi1383/simple-store/types"
	bolt "go.etcd.io/bbolt"
)

var (
	cfg types.Config
	db  *bolt.DB
)

func init() {
	cfg = config.GetConfig()
	var err error // No new variable should be initalized there
	db, err = bolt.Open(fmt.Sprintf("%v/store.db", cfg.BasePath), 0666, nil)
	if err != nil {
		panic(err)
	}
}

func SaveFile(dir, filename string, source io.Reader) error {
	fullPath := fmt.Sprintf("%v/%v", cfg.BasePath, dir)
	internalFilename := fmt.Sprintf("%v/%v", dir, filename)
	err := pathhelper.CreatePath(fullPath)
	if err != nil {
		return err
	}
	chunckID := 0
	for {
		dstFileName := fmt.Sprintf("%v__%v", internalFilename, chunckID)
		if _, err := os.Stat(dstFileName); !os.IsNotExist(err) {
			return err
		}
		dst, err := os.Create(dstFileName)
		if err != nil {
			return err
		}
		defer dst.Close()
		written, err := io.CopyN(dst, source, cfg.ChunkSize)
		if err == io.EOF {
			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucket([]byte(internalFilename))
				if err != nil {
					return err
				}
				return b.Put([]byte(dstFileName), typeconverters.Int64ToBytes(written))
			})
			if err != nil {
				return err
			}
			return nil
		} else if err != nil {
			return err
		} else {
			chunckID++
		}
	}
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
