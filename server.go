package main

import (
	"github.com/labstack/echo/v4"

	// "github.com/labstack/echo/v4/middleware"
	//"net/http"
	handlers "github.com/minhtungo/minh-file-api"
	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func main() {
	// create a new echo instance
	sh = shell.NewShell("localhost:8080")
	e := echo.New()
	// e.Use(middleware.CORS())

	// Routes
	e.GET("/", handlers.Hello)
	e.GET("/get", handlers.GetData)

	e.POST("/add", handlers.AddData)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
