package migrations

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Up_1730050275304(ctx context.Context, db *mongo.Database, client *mongo.Client) {
// Write your migration here
db.Collection("users").UpdateMany(ctx, map[string]interface{}{}, map[string]interface{}{
	"$rename": map[string]interface{}{
		"name": "full_name",
	},
})
}

func Down_1730050275304(ctx context.Context, db *mongo.Database, client *mongo.Client) {
// Write your rollback here
db.Collection("users").UpdateMany(ctx, map[string]interface{}{}, map[string]interface{}{
	"$rename": map[string]interface{}{
		"full_name": "name",
	},
})
}