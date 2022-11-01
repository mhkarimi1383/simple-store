package api

import (
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mhkarimi1383/simple-store/api/handlers"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Simple Store
// @version 1.0
// @description This is a simple api for storing files.
// @contact.email info@karimi.dev
// @BasePath /
func Serve(listenAddress string) {
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
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(listenAddress))
}
