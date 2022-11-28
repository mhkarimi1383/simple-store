package filemanager

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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
	fullFilename := fmt.Sprintf("%v/%v", fullPath, filename)
	err := pathhelper.CreatePath(fullPath)
	if err != nil {
		return err
	}
	var chunckID int64 = 0
	for {
		dstFileName := fmt.Sprintf("%v__%v", fullFilename, chunckID)
		if _, err := os.Stat(dstFileName); !os.IsNotExist(err) {
			return err
		}
		dst, err := os.Create(dstFileName)
		if err != nil {
			return err
		}
		defer dst.Close()
		written, err := io.CopyN(dst, source, cfg.ChunkSize)
		if err == io.EOF || err == nil {
			fileFinished := err == io.EOF
			err := db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucket([]byte(internalFilename))
				if err != nil {
					return err
				}
				return b.Put(typeconverters.Int64ToBytes(chunckID), typeconverters.Int64ToBytes(written))
			})
			if err != nil {
				return err
			}
			if fileFinished {
				return nil
			} else {
				chunckID++
			}
		} else if err != nil {
			return err
		}
	}
}

func RemoveFile(dir, filename string) error {
	internalFilename := fmt.Sprintf("%v/%v", dir, filename)
	return db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internalFilename))
		return b.ForEach(func(k, _ []byte) error {
			err := os.Remove(fmt.Sprintf("%v/%v__%v", cfg.BasePath, internalFilename, string(k)))
			if err != nil {
				return err
			}
			return nil
		})
	})
}

func GetFile(dir, filename string) (io.Reader, error) {
	internalFilename := fmt.Sprintf("%v/%v", dir, filename)
	content := []byte{}
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internalFilename))
		return b.ForEach(func(k, _ []byte) error {
			tmpContent, err := ioutil.ReadFile(fmt.Sprintf("%v/%v__%v", cfg.BasePath, internalFilename, string(k)))
			if err != nil {
				return err
			}
			content = append(content, tmpContent...)
			return nil
		})
	})
	return bytes.NewReader(content), err
}
