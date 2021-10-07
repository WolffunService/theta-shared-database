package util

import "go.mongodb.org/mongo-driver/bson"

func BsonIncrease(b bson.D,key string, value int) bson.D {
	if b == nil {b = bson.D{}}
	return append(b, bson.E{Key:"$inc", Value: bson.M{key:value}})
}
func BsonSet(b bson.D,key string, value interface{}) bson.D {
	if b == nil {b = bson.D{}}
	return append(b, bson.E{Key:"$set", Value: bson.M{key:value}})
}

func BsonPush(b bson.D,key string, value interface{}) bson.D {
	if b == nil {b = bson.D{}}
	return append(b, bson.E{Key:"$push", Value: bson.M{key:value}})
}

func BsonPushMultiArr(b bson.D,key string, value interface{}) bson.D {
	if b == nil {b = bson.D{}}
	return append(b, bson.E{Key:"$push", Value: bson.E{Key:"$each", Value: bson.M{key:value}}})
}
