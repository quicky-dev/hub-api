// Package db relates to database collections concerning the metadata related to script usage of the Hub API
package db

import (
	"context"
	"errors"
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

// SaveScript creates a new empty Script, population of Script attributes will have to be
// done seperately.
func (this *Script) SaveScript() error {
	insertion, err := script.InsertOne(context.TODO(), this)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	ID, status := insertion.InsertedID.(primitive.ObjectID)

	if status == false {
		return errors.New("Couldn't save the script.")
	}

	this.ID = thisID
	return nil

}

// GetScriptByID returns a specific script obj by id
func GetScriptByID(scriptID string) (*Script, error) {
	foundScript := new(Script)
	filter := bson.M{"_id": scriptID}
	err := script.FindOne(context.TODO(), filter).Decode(&foundScript)

	if err != nil {
		return foundScript, err
	}

	return foundScript, err
}

// ReturnScripts returns a list of all Scripts
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

// UpdateScript  updates a script associated with input id primitive. Returns true
// if update was successful, err otherwise.
func UpdateScriptComment(objectID string, comment Comment) (bool, error) {
	// get id of comment to update
	objectIDHex, err := util.GetObjectIDFromString(objectID)

	// TODO: ??
	// update := bson.D{
	// 	{"$inc", bson.D{
	// 		{"Comment", 1},
	// 	}},
	// }

	// update comment
	filter_comment := bson.D{{"Comments", "ID"}}

	// get scripts
	scripts, err := ReturnScripts()

	// loop through scripts
	for _, script := range scripts {
		for _, comment := range script.Comments {
			if comment.ID == objectIDHex {
				// update comment
				updateResult, err := collection.UpdateOne(context.TODO(), filter_comment, update)
				if err != nil {
					log.Fatal(err)
					return err
				}
			}

		}

	}
	return true, nil
}
