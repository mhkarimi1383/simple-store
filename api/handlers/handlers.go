package handlers

import (
	"net/http"
	"net/url"

	"github.com/gabriel-vasile/mimetype"
	"github.com/labstack/echo/v4"
	_ "github.com/mhkarimi1383/simple-store/api/docs"
	"github.com/mhkarimi1383/simple-store/internal/config"
	"github.com/mhkarimi1383/simple-store/internal/filemanager"
	"github.com/mhkarimi1383/simple-store/types"
)

var cfg types.Config

func init() {
	cfg = config.GetConfig()
}

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
	dir, err := url.QueryUnescape(c.Param("dir"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.HttpResponse{
			Error:   true,
			Message: "unable to unscape dir parameter",
			Details: &[]string{err.Error()},
		})
	}
	filename, err := url.QueryUnescape(c.Param("filename"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.HttpResponse{
			Error:   true,
			Message: "unable to unscape filename parameter",
			Details: &[]string{err.Error()},
		})
	}

	file, err := c.FormFile("data")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.HttpResponse{
			Error:   true,
			Message: "unable to open given file",
			Details: &[]string{err.Error()},
		})
	}
	err = filemanager.SaveFile(dir, filename, src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.HttpResponse{
			Error:   true,
			Message: "unable to save given file",
			Details: &[]string{err.Error()},
		})
	}
	return c.JSON(http.StatusCreated, types.HttpResponse{
		Error:   false,
		Message: "file saved",
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
	dir, err := url.QueryUnescape(c.Param("dir"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.HttpResponse{
			Error:   true,
			Message: "unable to unscape dir parameter",
			Details: &[]string{err.Error()},
		})
	}
	filename, err := url.QueryUnescape(c.Param("filename"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, types.HttpResponse{
			Error:   true,
			Message: "unable to unscape filename parameter",
			Details: &[]string{err.Error()},
		})
	}
	reader, err := filemanager.GetFile(dir, filename)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.HttpResponse{
			Error:   true,
			Message: "unable to read given file",
			Details: &[]string{err.Error()},
		})
	}
	mtype, err := mimetype.DetectReader(reader)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.HttpResponse{
			Error:   true,
			Message: "unable to detect mime type",
			Details: &[]string{err.Error()},
		})
	}
	return c.Stream(http.StatusOK, mtype.String(), reader)
}
