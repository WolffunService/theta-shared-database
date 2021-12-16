package main

import (
	"context"
	"fmt"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/useritems/useritemmodel"
	"github.com/WolffunGame/theta-shared-database/useritems/useritemstorage"
)

func main() {

	_, _, _ = mongodb.SetDefaultConfig(nil, nil)

	// oids := make([]primitive.ObjectID, 2)
	// objID1, _ := primitive.ObjectIDFromHex("614c85d856572a957b82f533")
	// objID2, _ := primitive.ObjectIDFromHex("6128e94fc92089fad4175c77")
	// oids[0] = objID1
	// oids[1] = objID2

	// col :=  mongodb.Coll(&usermodel.User{})
	// var users []usermodel.User
	// fmt.Println(col.FindByListID(oids, &users))
	// fmt.Println(len(users))
	//err := useritemstorage.AddAvatar(context.Background(), "61b9e6ed704751f713494eff", 1)
	//err := useritemstorage.AddListAvatar(context.Background(), "61b9e6ed704751f713494eff", 2, 3)
	err := useritemstorage.UpdateNewItem(context.Background(), "61b9e6ed704751f713494eff", 2, useritemmodel.AVATAR)
	fmt.Println(err)
}
