package main

import (
	mv "final_project/middleware"
	"final_project/route"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// log.Fatalf("config: %v \n", config.Config)

	e := echo.New()
	// Middleware
	e.Use(middleware.Recover())
	e.Use(mv.SimpleLogger())
	route.All(e)
	log.Println(e.Start(":9090"))

}
