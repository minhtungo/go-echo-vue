package main

import (
	"github.com/labstack/echo/v4"

	// "github.com/labstack/echo/v4/middleware"
	//"net/http"
	
	handlers "github.com/minhtungo/minh-file-api"
)

func main() {
	// create a new echo instance
	e := echo.New()

	// Routes
	e.GET("/", handlers.Hello)
	e.GET("/add", handlers.GetData)

	e.POST("/add", handlers.AddData)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

