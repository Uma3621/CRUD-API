package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// type Address struct {
// 	State   string `json:"state" bson:"state"`
// 	City    string `json:"city" bson:"city"`
// 	Pincode int    `json:"pincode" bson:"picode"`
// }
type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"user_name"`
	Age   int                `json:"age" bson:"user_age"`
	Email string             `json:"email" bson:"user_email"`
	//Address Address `json:"address" bson:"user_address"`
}
