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
func SaveScript() (string, error) {
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
func ReturnScripts() ([]*Script, error) {
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

	return results, nil
}

func UpdateScriptComment(objectID string, comment Comment) (bool, error) {
	// get id of comment to update
	objectIDHex, err := util.GetObjectIDFromString(objectID)

	update := bson.D{
		{"$inc", bson.D{
			{"Comment", 1},
		}},
	}

	// update comment
	filter_comment := bson.D{{"Comments", "ID"}}

	// get scripts
	scripts, err := ReturnScripts()

	// loop through scripts
	for _, script := range scripts {
		for _, comment := range Comments {
			if comment.ID == objectIDHex {
				// update comment
				updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	
}

// DeleteScript will delete the script associated with its id
func DropScript(objectID string) (bool, error) {
	objectIDHex, err := util.GetObjectIDFromString(objectID)

	if err != nil {
		return false, err
	}

	filter := bson.M{{"_id": objectIDHex}}

	_, err = script.DeleteOne(contex.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
