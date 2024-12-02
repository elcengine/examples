package seeds

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func Up_1730050801446(ctx context.Context, db *mongo.Database, client *mongo.Client) {
// Write your seed here
db.Collection("users").InsertOne(ctx, map[string]interface{}{
	"name": "Laksara2",
	"email": "laksara2@gmail.com",
})
}

func Down_1730050801446(ctx context.Context, db *mongo.Database, client *mongo.Client) {
// Write your rollback here
db.Collection("users").DeleteOne(ctx, map[string]interface{}{
	"email": "laksara@gmail.com",
})
}