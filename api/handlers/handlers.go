package handlers

import (
	"github.com/labstack/echo/v4"
	_ "github.com/mhkarimi1383/simple-store/api/docs"
	"github.com/mhkarimi1383/simple-store/internal/config"
	"github.com/mhkarimi1383/simple-store/types"
	"io"
	"net/http"
	"net/url"
	"os"
)

var cfg types.Config

// UploadFile
// @Summary Upload file
// @Description Upload file
// @ID file.upload
// @Accept  multipart/form-data
// @Param   data formData file true  "file to upload"
// @Param   dir path string true  "directory for file"
// @Param   filename path string true  "name for file"
// @Success 200 {string} string "ok"
// @Router /{dir}/{filename} [put]
func UploadFile(c echo.Context) error {
	if (cfg == types.Config{}) {
		cfg = config.GetConfig()
	}
	dir, err := url.QueryUnescape(c.Param("dir"))
	if err != nil {
		return err
	}
	dir = cfg.BasePath + "/" + dir
	filename, err := url.QueryUnescape(c.Param("filename"))
	if err != nil {
		return err
	}

	file, err := c.FormFile("data")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}

	dst, err := os.Create(dir + "/" + filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "File saved ;)",
	})
}

// DownloadFile
// @Summary Download file
// @Description Upload file
// @ID file.download
// @Param   dir path string true  "directory for file"
// @Param   filename path string true  "name for file"
// @Success 200 {string} string "ok"
// @Router /{dir}/{filename} [get]
func DownloadFile(c echo.Context) error {
	if (cfg == types.Config{}) {
		cfg = config.GetConfig()
	}
	dir, err := url.QueryUnescape(c.Param("dir"))
	if err != nil {
		return err
	}
	dir = cfg.BasePath + "/" + dir
	filename, err := url.QueryUnescape(c.Param("filename"))
	if err != nil {
		return err
	}
	return c.File(dir + "/" + filename)
}
