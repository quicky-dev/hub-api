// Package db relates to database collections concerning the metadata related to script usage of the Hub API
package db

import (
	"context"
	"fmt"
	"log"

	"github.com/quicky-dev/hub-api/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var script = util.MongoClient.Database("Hub").Collection("Script")

type Script struct {
	ID        primitive.ObjectID `bson: "_id, omitempty"`
	Username  string
	Auth_id   string
	Script_id string
	Ratings   primitive.ObjectID `bson: "_id, omitempty"`
	Comments  primitive.ObjectID `bson: "_id, omitempty"`
}

// CreateScript creates a new empty Script
func (this *Script) SaveScript() (string, error) {
	insertion, err := script.InsertOne(context.TODO(), this)

	if err != nil {
		log.Printf(err.Error())
		log.Printf("Could not save Script struct to collection")
	}

	thisID := util.GetObjectIdFromInsertion(insertion.InsertedID)

	if thisID == "" {
		return "", fmt.Errorf("Couldn't obtain the ObjectID for this: %s", this.Username)
	}

	return thisID, nil

}

// GetScripts returns a list of all Scripts
func (this *Script) GetScripts() []Script {

}

func (this *Script) DeleteScript(id int64) string {
	// delete script
}
