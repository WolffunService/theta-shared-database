package useritemstorage

import (
	"context"
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
	filter := bson.M{field.ID: objectUserId}
	// update
	update := bson.D{}
	update = util.BsonPush(update, "avatars", avatarId)

	return updateOneUserItems(ctx, filter, update)
}

func AddListAvatar(ctx context.Context, userId string, avatarIds []int) error {
	objectUserId := util.ObjectIDFromHex(userId)
	filter := bson.M{field.ID: objectUserId}
	// update
	update := bson.D{}
	update = util.BsonPushMultiArr(update, "avatars", avatarIds)

	return updateOneUserItems(ctx, filter, update)
}

func updateOneUserItems(ctx context.Context, filter, update interface{}, opts ...*options.UpdateOptions) error {
	col := mongodb.Coll(&useritemmodel.UserItems{}, mongodb.WriteConcernW1())
	opts = append(opts, mongodb.UpsertTrueOption())
	_, err := col.UpdateOne(ctx, filter, update, opts...)
	return err
}
