package main

import (
	"fmt"
	"github.com/WolffunGame/theta-shared-database/database/mongodb"
	"github.com/WolffunGame/theta-shared-database/user/usermodel"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main()  {

	_ ,_ ,_ = mongodb.SetDefaultConfig(nil,nil)

	oids := make([]primitive.ObjectID, 2)
	objID1, _ := primitive.ObjectIDFromHex("614c85d856572a957b82f533")
	objID2, _ := primitive.ObjectIDFromHex("6128e94fc92089fad4175c77")
	oids[0] = objID1
	oids[1] = objID2

	col :=  mongodb.Coll(&usermodel.User{})
	var users []usermodel.User
	fmt.Println(col.FindByListID(oids, &users))
	fmt.Println(len(users))
}

