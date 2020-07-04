package main

import (
	"github.com/Tak1za/backer/config"
	"github.com/Tak1za/backer/handlers"
	"github.com/Tak1za/backer/middlewares"
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
	router.POST("/api/follow", handlers.FollowUser)

	router.Run()
}
