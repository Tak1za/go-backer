package main

import (
	"github.com/Tak1za/ivar/config"
	"github.com/Tak1za/ivar/handlers"
	"github.com/Tak1za/ivar/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	driver, err := config.InitDriver()
	if err != nil {
		panic(err)
	}

	defer driver.Close()

	env := &config.Env{Driver: driver}

	router := gin.Default()

	router.Use(
		middlewares.CorsMiddleware(),
		middlewares.DriverMiddleware(env),
	)

	router.POST("/api/users", handlers.CreateUser)

	router.Run()
}
