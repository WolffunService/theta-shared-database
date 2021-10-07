package useritemmodel

import "github.com/WolffunGame/theta-shared-database/database/mongodb"

func (UserItems) CollectionName() string {
	return "UserItems"
}

type UserItems struct {
	mongodb.IDIntField `bson:",inline"`
	Avatars []int  `bson:"avatars" json:"avatars"`
}