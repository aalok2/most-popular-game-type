package tests

import (
    "context"
    "testing"
    "multiplayer-mode-usage/cache"
    "multiplayer-mode-usage/db"
    "multiplayer-mode-usage/service"
    "go.mongodb.org/mongo-driver/bson"
)

func TestIntegrationPopularMode(t *testing.T) {
    db.InitializeMongoDB("mongodb://localhost:27017", "test_db")

    cache.InitRedis("localhost:6379")

    coll := db.GetDB().Collection("mode_usage")
    coll.Drop(context.TODO())
    testDoc := bson.D{{"area_code", "123"}, {"mode_name", "Battle Royale"}, {"player_count", 200}}
    _, err := coll.InsertOne(context.TODO(), testDoc)
    if err != nil {
        t.Fatalf("Failed to insert test document: %v", err)
    }
    mode, count, err := service.GetPopularMode("123")
    if err != nil {
        t.Fatalf("Failed to get popular mode: %v", err)
    }
    if mode != "Battle Royale" {
        t.Errorf("Expected mode 'Battle Royale', got %s", mode)
    }
    if count != 200 {
        t.Errorf("Expected player count 200, got %d", count)
    }

    coll.DeleteMany(context.TODO(), bson.M{"area_code": "123"})
}
