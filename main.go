package main

import (
	_ "perqara_api/docs"
	"perqara_api/models"
	route "perqara_api/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Perqara API
// @version 1.0
// @description This is a sample API for managing users.
// @host localhost:8080
// @BasePath /

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Initialize database
	models.InitDB()

	// Routes
	route.InitRoutes(e)

	// Swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
