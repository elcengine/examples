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
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	Name             string             `json:"name,omitempty" bson:"name,omitempty"`
	Email            string             `json:"email,omitempty" bson:"email,omitempty"`
	Age              int                `json:"age,omitempty" bson:"age,omitempty"`
	Password         string             `json:"password,omitempty" bson:"password,omitempty"`
	Organizations    []string           `json:"organizations,omitempty" bson:"organizations"`
	Verified         bool               `json:"verified,omitempty" bson:"verified"`
	VerificationCode *string            `json:"verification_code,omitempty" bson:"verification_code,omitempty"`
	Role             UserRole           `json:"role,omitempty" bson:"role,omitempty"`
	CreatedAt        *time.Time         `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt        *time.Time         `json:"updated_at,omitempty" bson:"updated_at"`
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
