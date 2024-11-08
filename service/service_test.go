package service

import (
    "testing"
    "multiplayer-mode-usage/cache"
    "multiplayer-mode-usage/db"
    "multiplayer-mode-usage/service"
    "github.com/stretchr/testify/assert"
)

func TestGetPopularMode_CacheHit(t *testing.T) {
    cache.InitRedis("localhost:6379")
    cache.SetPopularModeCache("CA", "Battle Royale", 300, 5*time.Minute)

    mode, count, err := service.GetPopularMode("CA")
    assert.NoError(t, err)
    assert.Equal(t, "Battle Royale", mode)
    assert.Equal(t, int32(300), count)
}

func TestGetPopularMode_CacheMiss_DBHit(t *testing.T) {
    db.InitializeMongoDB("mongodb://localhost:27017", "test_db")
    testDB := db.GetDB()
    collection := testDB.Collection("mode_usage")
    collection.InsertOne(context.TODO(), bson.M{
        "area_code":    "TX",
        "mode_name":    "Deathmatch",
        "player_count": 100,
    })

    mode, count, err := service.GetPopularMode("TX")
    assert.NoError(t, err)
    assert.Equal(t, "Deathmatch", mode)
    assert.Equal(t, int32(100), count)

    collection.Drop(context.TODO())
}
