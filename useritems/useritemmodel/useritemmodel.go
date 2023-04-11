package useritemmodel

import "github.com/WolffunService/theta-shared-database/database/mongodb"

func (UserItems) CollectionName() string {
	return "UserItems"
}

type UserItems struct {
	mongodb.IDField `bson:",inline"`
	Avatars         map[int]ItemModel `bson:"avatars" json:"avatars"`
}

type ItemModel struct {
	ItemId   int      `bson:"itemId" json:"itemId"`
	ItemType ItemType `bson:"itemType" json:"itemType"`
	//Amount   int      `bson:"amount,omitempty" json:"amount,omitempty"`
	NewItem bool `bson:"newItem" json:"newItem"`
}

func NewItems(itemType ItemType, itemId int, isNew bool) *ItemModel {
	return &ItemModel{
		itemId, itemType, isNew,
	}
}

type ItemType int

const (
	AVATAR ItemType = iota
)
