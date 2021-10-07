package useritemstorage

import (
	"context"
	"fmt"
	"github.com/WolffunGame/theta-shared-database/common/util"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/useritems/useritemmodel"
)

func FindUserItemsById(ctx context.Context, userId string) (*useritemmodel.UserItems, error) {
	userItems := &useritemmodel.UserItems{}
	userObjectId := util.ObjectIDFromHex(userId)
	col := mongodb.Coll(userItems)
	err := col.FindByIDWithCtx(ctx, userObjectId, userItems)
	if err != nil {
		return userItems, err
	}
	fmt.Printf("Found a single document: %+v\n", userItems)
	return userItems, nil
}
