// Package api api starts from here
package api

import (
	"net/http"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mhkarimi1383/simple-store/api/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// Serve starts http server using echo framework
// @title Simple Store
// @version 1.0
// @description This is a simple api for storing files.
// @contact.email info@karimi.dev
// @BasePath /
func Serve(listenAddress string, enableSwagger bool) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	p := prometheus.NewPrometheus("simple_store", nil)
	p.Use(e)

	// Routes
	e.PUT("/:dir/:filename", handlers.UploadFile)
	e.GET("/:dir/:filename", handlers.DownloadFile)
	e.DELETE("/:dir/:filename", handlers.DeleteFile)
	if enableSwagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
		e.GET("/swagger", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "/swagger/")
		})
	}

	// Start server
	e.Logger.Fatal(e.Start(listenAddress))
}
