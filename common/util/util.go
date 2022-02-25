package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ObjectIDFromHex(s string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return objID
}

func GetObjectIDFromHex(s string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objID, nil
}