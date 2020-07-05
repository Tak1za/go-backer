package handlers

import (
	"errors"
	"log"

	"github.com/Tak1za/go-backer/config"
	"github.com/gin-gonic/gin"
)

func getDriver(c *gin.Context) (*config.Env, error) {
	driver, ok := c.MustGet("driver").(*config.Env)
	if !ok {
		errMessage := "Error getting Neo4j driver"
		log.Fatalln(errMessage)
		return nil, errors.New(errMessage)
	}

	return driver, nil
}
