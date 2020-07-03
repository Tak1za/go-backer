package handlers

import (
	"log"
	"net/http"

	"github.com/Tak1za/ivar/config"
	"github.com/Tak1za/ivar/models"
	"github.com/Tak1za/ivar/service"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	driver, ok := c.MustGet("driver").(*config.Env)
	if !ok {
		errMessage := "Error getting Neo4j driver"
		log.Fatalln(errMessage)
		c.JSON(http.StatusInternalServerError, gin.H{"error": errMessage})
		return
	}

	var createUserRequest models.CreateUserRequest
	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		log.Println("Request is not valid")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	session, err := driver.GetWriteSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer session.Close()

	ce := make(chan error)

	go service.CreateUser(session, createUserRequest, ce)

	if err = <-ce; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
