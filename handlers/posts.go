package handlers

import (
	"log"
	"net/http"

	"github.com/Tak1za/go-backer/models"
	"github.com/Tak1za/go-backer/service"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var createPostRequest models.CreatePostRequest
	if err := c.ShouldBindJSON(&createPostRequest); err != nil {
		log.Println("Not a valid request: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not a valid request"})
		return
	}

	driver, err := getDriver(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, err := driver.GetWriteSession()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Something unexpected happened at the server"})
		return
	}

	defer session.Close()

	ce := make(chan error)

	go service.CreatePost(session, createPostRequest, ce)

	if err = <-ce; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
