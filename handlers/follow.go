package handlers

import (
	"log"
	"net/http"

	"github.com/Tak1za/go-backer/models"
	"github.com/Tak1za/go-backer/service"
	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	driver, err := getDriver(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var followRequest models.FollowRequest
	if err := c.ShouldBindJSON(&followRequest); err != nil {
		log.Println("Not a valid request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid request"})
		return
	}

	session, err := driver.GetWriteSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something unexpected happened at the server"})
		return
	}

	defer session.Close()

	ce := make(chan error)

	go service.FollowUser(session, followRequest, ce)

	if err = <-ce; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to follow user"})
		return
	}

	c.Status(http.StatusCreated)
}
