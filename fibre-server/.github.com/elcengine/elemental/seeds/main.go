package main 

import (
	"context"
	"strings"
	"strconv"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/samber/lo"
	"github.com/elcengine/elemental/connection"
	"github.com/elcengine/elemental/database/seeds"
)

 func extractTimestamp(fileName string) int64 {
		fileName = strings.TrimSuffix(fileName, ".go")
		parts := strings.Split(fileName, "_")
		timestampStr := parts[len(parts)-1]
		return lo.Must(strconv.ParseInt(timestampStr, 10, 64))
}

func main() {
	client := e_connection.ConnectURI("mongodb+srv://akalankaperera128:pFAnQVXE6vrbcXNk@default.ynr156r.mongodb.net/elemental")
	db := e_connection.UseDefault()
	go db.Collection("changelog").Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: map[string]interface{}{"type": 1},
	})
	files := []string{"test_1730050801446"}
	functions := []func(context.Context, *mongo.Database, *mongo.Client){seeds.Up_1730050801446}
	ctx := context.Background()
	for i, f := range files {
		functionToRun := functions[i]
		var entry = map[string]interface{}{}
		db.Collection("changelog").FindOne(ctx, map[string]interface{}{"name": f, "type": "seed"}).Decode(&entry)
		if false {
			if entry["name"] != nil {
				functionToRun(ctx, db, &client)
				db.Collection("changelog").DeleteOne(ctx, map[string]interface{}{"name": f})
			}
		} else {
			if entry["name"] == nil {
				functionToRun(ctx, db, &client)
				db.Collection("changelog").InsertOne(ctx, map[string]interface{}{
					"name": f,
					"type": "seed",
					"created_at": time.Now(),
				})
			} 
		}
	}
}
