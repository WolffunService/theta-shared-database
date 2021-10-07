package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDFromHex(s string) (primitive.ObjectID, error) {
	objID, err := primitive.ObjectIDFromHex(s)
	if err != nil {
		panic(err)
	}
	return objID, nil
}
