package util

import "go.mongodb.org/mongo-driver/bson"

func BsonAdd(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: value})
}

func BsonIncrease(b bson.D, key string, value int) bson.D {
	return append(b, bson.E{Key: "$inc", Value: bson.M{key: value}})
}

func BsonIncrease64(b bson.D, key string, value int64) bson.D {
	return append(b, bson.E{Key: "$inc", Value: bson.M{key: value}})
}

func BsonSet(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: "$set", Value: bson.M{key: value}})
}

func BsonUnSet(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: "$unset", Value: bson.M{key: value}})
}

func BsonEqual(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$eq": value}})
}

func BsonNotEqual(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$ne": value}})
}

func BsonGreaterThan(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$gt": value}})
}

func BsonGreaterThanEqual(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$gte": value}})
}

func BsonLessThan(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$lt": value}})
}

func BsonLessThanEqual(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$lte": value}})
}

func BsonIn(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$in": value}})
}
func BsonNotIn(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: key, Value: bson.M{"$nin": value}})
}

func BsonPush(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: "$push", Value: bson.M{key: value}})
}

func BsonPushMultiArr(b bson.D, key string, value interface{}) bson.D {
	return append(b, bson.E{Key: "$push", Value: bson.M{key: bson.M{"$each": value}}})
}
