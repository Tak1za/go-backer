package service

import (
	"log"

	"github.com/Tak1za/go-backer/models"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CreateUser(session neo4j.Session, newUserRequest models.CreateUserRequest, ce chan error) {
	query := "CREATE (n:User {name: $name, email: $email, gender: $gender, image: $image})"
	result, err := session.Run(query, map[string]interface{}{
		"name":   newUserRequest.Name,
		"email":  newUserRequest.Email,
		"gender": newUserRequest.Gender,
		"image":  newUserRequest.Image,
	})

	if err != nil {
		log.Println("Error running create user query", err.Error())
		ce <- err
		return
	}

	if err = result.Err(); err != nil {
		log.Println("Error creating user", err.Error())
		ce <- err
		return
	}

	log.Println(result.Summary())
	log.Println("User created")
	ce <- nil
	return
}
