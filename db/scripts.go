// Package db relates to database collections concerning the metadata related to script usage of the Hub API
package db

import (
	"context"
	"fmt"
	"log"

	"github.com/quicky-dev/hub-api/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var script = util.MongoClient.Database("Hub").Collection("Script")

type Script struct {
	ID        primitive.ObjectID `bson: "_id, omitempty"`
	Username  string
	Auth_id   string
	Script_id string
	Ratings   []primitive.ObjectID
	Comments  []primitive.ObjectID
}

// CreateScript creates a new empty Script, population of Script attributes will have to be
// done seperately.
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
func (this *Script) GetScripts() error {
	findOptions := options.Find()
	empty_filter := bson.D{{}}

	var results []*Script

	find, err := script.Find(context.TODO(), empty_filter, findOptions)

	// if multiple documents are found, a cursor will be returned
	// iterate through cursor
	for find.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Script
		err := find.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (this *Script) DeleteScript(id int64) string {
	// delete script
}
