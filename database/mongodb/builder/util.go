package builder

import (
	"go.mongodb.org/mongo-driver/bson"
	"wolffundb/database/mongodb/utils"
)

// appendIfHasVal append key and val to map if value is not empty.
func appendIfHasVal(m bson.M, key string, val interface{}) {
	if !utils.IsNil(val) {
		m[key] = val
	}
}
