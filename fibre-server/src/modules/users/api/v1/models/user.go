package models

import (
	"reflect"
	"time"

	"github.com/elcengine/elemental/core"
	"github.com/samber/lo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRole string

const (
	Admin UserRole = "admin"
	Guest UserRole = "guest"
)

type User struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	Name             string             `json:"name" bson:"name,omitempty"`
	Email            string             `json:"email" bson:"email,omitempty"`
	Age              int                `json:"age" bson:"age,omitempty"`
	Password         string             `json:"password" bson:"password,omitempty"`
	Organizations    []string           `json:"organizations" bson:"organizations"`
	Verified         bool               `json:"verified" bson:"verified"`
	VerificationCode *string            `json:"verification_code" bson:"verification_code,omitempty"`
	Role             UserRole           `json:"role" bson:"role,omitempty"`
	CreatedAt        time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at" bson:"updated_at"`
}

var UserModel = elemental.NewModel[User]("User", elemental.NewSchema(map[string]elemental.Field{
	"Name": {
		Type:     reflect.String,
		Required: true,
	},
	"Email": {
		Type:     reflect.String,
		Required: true,
		Index: options.IndexOptions{
			Unique: lo.ToPtr(true),
		},
	},
	"Age": {
		Type: reflect.Int,
	},
	"Password": {
		Type:     reflect.String,
		Required: true,
	},
	"Organizations": {
		Type:    reflect.Slice,
		Default: []string{},
	},
	"Verified": {
		Type:    reflect.Bool,
		Default: false,
	},
	"VerificationCode": {
		Type: reflect.String,
	},
	"Role": {
		Type:    reflect.String,
		Default: Guest,
	},
}, elemental.SchemaOptions{
	Collection: "users",
}))

func (u User) Secure() User {
	u.Password = ""
	return u
}
