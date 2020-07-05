package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/Tak1za/go-backer/models"
	"github.com/neo4j/neo4j-go-driver/neo4j"
)

func CreatePost(session neo4j.Session, newPostRequest models.CreatePostRequest, ce chan error) {
	query := `
		MATCH (author:User)
		WHERE author.email = $author
		MERGE (author)-[:POSTED]->(post:Post {content: $content, description: $description, postedAt: $timestamp})
		RETURN author, post
	`

	result, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error){
		res, err := transaction.Run(
			query, map[string]interface{}{
				"author":      newPostRequest.Author,
				"content":     newPostRequest.Content,
				"description": newPostRequest.Description,
				"timestamp":   neo4j.LocalDateTimeOf(time.Now()),
			})

		//Run failure
		if err != nil {
			log.Println(err.Error())
			return nil, errors.New("Failed to create Post")
		}

		if res.Next(){
			if _, ok := res.Record().Get("author"); ok {
				if _, ok := res.Record().Get("post"); ok {
					summary, err := res.Consume()
					if err != nil {
						log.Println(err.Error())
						return nil, errors.New("Failed to create Post")
					}

					return summary, nil
				}
			}
		}
		
		err = errors.New(fmt.Sprintf("Author %s does not exist", newPostRequest.Author))
		log.Println(err.Error())
		return nil, err
	})

	//Write failure
	if err != nil {
		ce <- err
		return
	}

	//Write success
	log.Println("Post created successfully: ", result)
	ce <- nil
	return
}