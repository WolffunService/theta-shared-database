package useritemstorage

import (
	"context"
	"fmt"

	"github.com/WolffunService/theta-shared-database/common/util"
	"github.com/WolffunService/theta-shared-database/database/mongodb"
	"github.com/WolffunService/theta-shared-database/useritems/useritemmodel"
)

func FindUserItemsById(ctx context.Context, userId string) (*useritemmodel.UserItems, error) {
	userItems := &useritemmodel.UserItems{}
	userObjectId := util.ObjectIDFromHex(userId)
	col := mongodb.CollRead(userItems)
	err := col.FindByIDWithCtx(ctx, userObjectId, userItems)
	if err != nil {
		userItems.ID = userObjectId
		return userItems, err
	}
	fmt.Printf("Found a single document: %+v\n", userItems)
	return userItems, nil
}
