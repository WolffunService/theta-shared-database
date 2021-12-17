package useritemstorage

import (
	"context"
	"fmt"
	"github.com/WolffunGame/theta-shared-database/common/util"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/database/mongodb/field"
	"github.com/WolffunGame/theta-shared-database/useritems/useritemmodel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateUserItems(ctx context.Context, userItems *useritemmodel.UserItems) error {
	col := mongodb.Coll(userItems)
	return col.UpdateWithCtx(ctx, userItems)
}

func AddAvatar(ctx context.Context, userId string, avatarId int) error {
	objectUserId := util.ObjectIDFromHex(userId)
	filter := util.BsonAdd(nil, field.ID, objectUserId)
	filter = util.BsonAdd(filter, fmt.Sprintf("avatars.%d", avatarId), bson.M{"$exists": false})
	// update
	update := bson.D{}
	update = util.BsonSet(update, fmt.Sprintf("avatars.%d", avatarId),
		useritemmodel.NewItems(useritemmodel.AVATAR, avatarId, true))

	return updateOneUserItems(ctx, filter, update)
}

func AddListAvatar(ctx context.Context, userId string, avatarIds ...int) error {
	if len(avatarIds) == 0 {
		return fmt.Errorf("avatarIds is empty")
	}
	objectUserId := util.ObjectIDFromHex(userId)
	filter := util.BsonAdd(nil, field.ID, objectUserId)
	// update
	update := bson.D{}
	for _, avatarId := range avatarIds {
		filter = util.BsonAdd(filter, fmt.Sprintf("avatars.%d", avatarId), bson.M{"$exists": false})
		update = util.BsonSet(update, fmt.Sprintf("avatars.%d", avatarId),
			useritemmodel.NewItems(useritemmodel.AVATAR, avatarId, true))
	}

	return updateOneUserItems(ctx, filter, update)
}

func UpdateNewItem(ctx context.Context, userId string, itemId int, itemType useritemmodel.ItemType) error {
	key := ""
	switch itemType {
	case useritemmodel.AVATAR:
		key = "avatars"
	default:
		return fmt.Errorf("itemType was not config %v", itemType)
	}

	objectUserId := util.ObjectIDFromHex(userId)
	filter := bson.M{field.ID: objectUserId, fmt.Sprintf("%s.%d", key, itemId): bson.M{"$exists": true}}
	update := util.BsonSet(nil, fmt.Sprintf("%s.%d.newItem", key, itemId), false)
	return updateOneUserItems(ctx, filter, update)
}

func updateOneUserItems(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) error {
	col := mongodb.Coll(&useritemmodel.UserItems{}, mongodb.WriteConcernW1())
	opts = append(opts, mongodb.UpsertTrueOption())
	_, err := col.UpdateOne(ctx, filter, update, opts...)
	return err
}
