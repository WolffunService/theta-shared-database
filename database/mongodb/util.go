package mongodb

import (
	"github.com/jinzhu/inflection"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
	"wolffundb/database/mongodb/utils"
)

// Coll return model's collection.
func Coll(m Model, opts ...*options.CollectionOptions) *Collection {

	if collGetter, ok := m.(CollectionGetter); ok {
		return collGetter.Collection()
	}
	return CollectionByName(CollName(m), opts...)
}

// CollName check if you provided collection name in your
// model, return it's name, otherwise guess model
// collection's name.
func CollName(m Model) string {
	if collNameGetter, ok := m.(CollectionNameGetter); ok {
		return collNameGetter.CollectionName()
	}
	name := reflect.TypeOf(m).Elem().Name()
	return inflection.Plural(utils.ToSnakeCase(name))
}

// UpsertTrueOption returns new instance of the
//UpdateOptions with upsert=true property.
func UpsertTrueOption() *options.UpdateOptions {
	upsert := true
	return &options.UpdateOptions{Upsert: &upsert}
}

