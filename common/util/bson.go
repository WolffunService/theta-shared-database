package util

import "go.mongodb.org/mongo-driver/bson"

func BsonAdd(b bson.D,key string, value interface{}) bson.D {
	return append(b, bson.E{Key:key, Value: value})
}

func BsonIncrease(b bson.D,key string, value int) bson.D {
	return append(b, bson.E{Key:"$inc", Value: bson.M{key:value}})
}

func BsonIncrease64(b bson.D,key string, value int64) bson.D {
	return append(b, bson.E{Key:"$inc", Value: bson.M{key:value}})
}

func BsonSet(b bson.D,key string, value interface{}) bson.D {
	return append(b, bson.E{Key:"$set", Value: bson.M{key:value}})
}

func BsonNotEqual(b bson.D,key string, value interface{}) bson.D {
	return append(b, bson.E{Key:key, Value: bson.M{"$ne":value}})
}

func BsonPush(b bson.D,key string, value interface{}) bson.D {
	return append(b, bson.E{Key:"$push", Value: bson.M{key:value}})
}

func BsonPushMultiArr(b bson.D,key string, value interface{}) bson.D {
	return append(b, bson.E{Key:"$push", Value: bson.E{Key:"$each", Value: bson.M{key:value}}})
}
