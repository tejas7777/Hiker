package main

import (
	"fmt"
	api "hiker/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	env "hiker/env"
	test "hiker/test"
)

func main() {
	fmt.Println("Hello")
	env.SetEnvVariables()
	test.TestConnection()
	//initialize echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.POST("/auth", api.Auth)
	e.GET("/check", api.Check)
	e.GET("/trails", api.Trails)
	// Start server
	//e.Logger.Fatal(e.StartTLS(":1323", "cert.pem", "key.pem"))
	e.Logger.Fatal(e.Start(":1323"))

	//Allow Cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

}
