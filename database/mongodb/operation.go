package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/WolffunGame/theta-shared-database/database/mongodb/field"
)

func create(ctx context.Context, c *Collection, model Model, opts ...*options.InsertOneOptions) error {
	// Call to saving hook
	if err := callToBeforeCreateHooks(model); err != nil {
		return err
	}

	res, err := c.InsertOne(ctx, model, opts...)

	if err != nil {
		return err
	}

	// Set new id
	model.SetID(res.InsertedID)

	return callToAfterCreateHooks(model)
}

func createMany(ctx context.Context, c *Collection, documents []interface{}, opts ...*options.InsertManyOptions) error {
	//TODO update check hook
	//documents := util.InterfaceSlice(models)
	_, err := c.InsertMany(ctx, documents, opts...)

	if err != nil {
		return err
	}
	return nil
}

func first(ctx context.Context, c *Collection, filter interface{}, model Model, opts ...*options.FindOneOptions) error {
	return c.FindOne(ctx, filter, opts...).Decode(model)
}

func firstAndUpdate(ctx context.Context, c *Collection, filter interface{},update interface{}, model Model, opts ...*options.FindOneAndUpdateOptions) error {
	return c.FindOneAndUpdate(ctx, filter,update, opts...).Decode(model)
}

func update(ctx context.Context, c *Collection, model Model, opts ...*options.UpdateOptions) error {
	// Call to saving hook
	if err := callToBeforeUpdateHooks(model); err != nil {
		return err
	}
	res, err := c.UpdateOne(ctx, bson.M{field.ID: model.GetID()}, bson.M{"$set": model}, opts...)

	if err != nil {
		return err
	}

	return callToAfterUpdateHooks(res, model)
}

func del(ctx context.Context, c *Collection, model Model) error {
	if err := callToBeforeDeleteHooks(model); err != nil {
		return err
	}
	res, err := c.DeleteOne(ctx, bson.M{field.ID: model.GetID()})
	if err != nil {
		return err
	}

	return callToAfterDeleteHooks(res, model)
}
func count(ctx context.Context, c *Collection,filter interface{}, opts ...*options.CountOptions) (int64,error) {
	count,err :=  c.CountDocuments(ctx,filter,opts...)
	return count,err
}