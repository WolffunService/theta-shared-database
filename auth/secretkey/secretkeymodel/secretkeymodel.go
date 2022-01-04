package secretkeymodel

import (
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"time"
)

func (SecretKey) CollectionName() string {
	return "SecretKey"
}

type SecretKey struct {
	mongodb.DefaultModel `json:",inline" bson:",inline"`
	Key                  			string 							`json:"key" bson:"key"`
	Secret                  		string 							`json:"secret" bson:"secret"`
	KeyType        					SecretKeyType    				`json:"keyType" bson:"keyType"`
	Expired 						time.Time 						`json:"expired" bson:"expired"`
	Status 							int 							`json:"status" bson:"status"`
}

type SecretKeyType int

const (
	KEYTYPE_QUANTUM					SecretKeyType = 1
)