package db

import (
	"context"
	"errors"
	"github.com/quicky-dev/auth-service/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

var comments = util.MongoClient.Database("Hub").Collection("Comments")

type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Content  string             `bson:"content"`
	Username string             `bson:"username"`
	AuthID   string             `bson:"authID"`
}

func GetCommentByID(commentID string) (*Comment, error) {
	foundComment := new(Comment)
	filter := bson.M{"_id": commentID}
	err := comments.FindOne(context.TODO(), filter).Decode(&foundComment)

	if err != nil {
		return foundComment, err
	}

	return foundComment, nil
}

func (this *Comment) Save() error {
	insertion, err := comments.InsertOne(context.TODO(), this)

	if err != nil {
		log.Println(err)
		return err
	}

	ID, status := insertion.InsertedID.(primitive.ObjectID)

	if status == false {
		return errors.New("Couldn't save the comment.")
	}

	this.ID = ID
	return nil
}
