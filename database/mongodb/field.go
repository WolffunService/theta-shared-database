package mongodb

import (
	"time"
)

// IDField struct contain model's ID field.
type IDField struct {
	ID interface{} `json:"id" bson:"_id,omitempty"`
}

// IDField struct contain model's ID field.
type IDIntField struct {
	ID int64 `json:"id" bson:"_id,omitempty" csv:"id"`
}

// DateFields struct contain `created_at` and `updated_at`
// fields that autofill on insert/update model.
type DateFields struct {
	CreatedAt time.Time `json:"created" bson:"created"`
	UpdatedAt time.Time `json:"lastModified" bson:"lastModified"`
}

// PrepareID method prepare id value to using it as id in filtering,...
// e.g convert hex-string id value to bson.ObjectId
//func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
//	if idStr, ok := id.(string); ok {
//		return primitive.ObjectIDFromHex(idStr)
//	}
//
//	// Otherwise id must be ObjectId
//	return id, nil
//}

// GetID method return model's id
func (f *IDIntField) GetID() interface{} {
	return f.ID
}

// SetID set id value of model's id field.
func (f *IDIntField) SetID(id interface{}) {
	f.ID = id.(int64)
}

// GetID method return model's id
func (f *IDField) GetID() interface{} {
	return f.ID
}

// SetID set id value of model's id field.
func (f *IDField) SetID(id interface{}) {
	f.ID = id
}

//--------------------------------
// DateField methods
//--------------------------------

// Creating hook used here to set `created_at` field
// value on inserting new model into database.
func (f *DateFields) Creating() error {
	//f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook used here to set `updated_at` field value
// on create/update model.
func (f *DateFields) Saving() error {
	//f.UpdatedAt = time.Now().UTC()
	return nil
}
