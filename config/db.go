package config

import (
	"log"
	"sync"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

type Env struct {
	neo4j.Driver
}

func InitDriver() (neo4j.Driver, error) {
	var once sync.Once
	var driver neo4j.Driver
	var driverErr error
	once.Do(
		func() {
			driver, driverErr = neo4j.NewDriver("bolt://localhost:7687", neo4j.BasicAuth("neo4j", "password", ""), func(config *neo4j.Config) {
				config.Encrypted = false
			})
			if driverErr != nil {
				log.Println("Error initializing Neo4j driver: ", driverErr)
			}
			log.Println("Initialized Neo4j driver.")
		})
	return driver, driverErr
}

func (env *Env) GetReadSession() (neo4j.Session, error) {
	sessionConfig := neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: "backer",
	}

	session, err := env.NewSession(sessionConfig)
	if err != nil {
		log.Println("Error creating read session", err)
		return nil, err
	}

	log.Println("Created read session")
	return session, nil
}

func (env *Env) GetWriteSession() (neo4j.Session, error) {
	sessionConfig := neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: "backer",
	}

	session, err := env.NewSession(sessionConfig)
	if err != nil {
		log.Println("Error creating write session", err)
		return nil, err
	}

	log.Println("Created write session")
	return session, nil
}
