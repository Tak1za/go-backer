package service

import (
	"errors"
	"log"

	"github.com/Tak1za/go-backer/models"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func FollowUser(session neo4j.Session, newFollowRequest models.FollowRequest, ce chan error) {
	query := `
		MATCH (follower:User)
		WHERE follower.email = $followerEmail
		MATCH (following:User)
		WHERE following.email = $followingEmail
		MERGE (follower)-[:FOLLOWS]->(following)
	`
	result, err := session.Run(query, map[string]interface{}{
		"followerEmail":  newFollowRequest.Follower,
		"followingEmail": newFollowRequest.Following,
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

	summary, err := result.Summary()
	if err != nil {
		log.Println("Error getting result summary", err.Error())
		ce <- err
		return
	}

	log.Println("Follow User Summary: ", summary)

	if summary.Counters().RelationshipsCreated() == 0 {
		log.Println("Relationship already exists")
		ce <- errors.New("Relationship already exists")
		return
	}

	log.Println("User followed successfully")
	ce <- nil
	return
}
